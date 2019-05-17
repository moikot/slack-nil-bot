package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/nlopes/slack"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	defaultPromAddr = ":9153"
)

var (
	connStatus = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "nil_bot_connection_status",
		Help: "Nil bot connection status",
	})
)

func getEnv(name string) string {
	v := os.Getenv(name)
	if v == "" {
		panic("missing required environment variable " + name)
	}
	return v
}

func getPromAddr() string {
	v := os.Getenv("PROM_ADDR")
	if v == "" {
		return defaultPromAddr
	}
	return v
}

func main() {
	token := getEnv("BOT_TOKEN")
	api := slack.New(token)
	rtm := api.NewRTM()

	go rtm.ManageConnection()

	go startHttp(getPromAddr())

Loop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.RTMError:
				fmt.Printf("error: %s\n", ev.Error())

			case *slack.InvalidAuthEvent:
				fmt.Printf("error: invalid authentication\n")
				break Loop

			case *slack.ConnectedEvent:
				connStatus.Set(1)

			case *slack.DisconnectedEvent:
				connStatus.Set(0)
				if ev.Intentional {
					break Loop
				}

			default:
			}
		}
	}
}

func startHttp(addr string) {
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		panic("unable to create HTTP server, error: " + err.Error())
	}
}
