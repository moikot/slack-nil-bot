package main

import (
	"fmt"
	"github.com/nlopes/slack"
	"os"
)

func getEnv(name string) string {
	v := os.Getenv(name)
	if v == "" {
		panic("missing required environment variable " + name)
	}
	return v
}

func main() {
	token := getEnv("BOT_TOKEN")
	api := slack.New(token)
	rtm := api.NewRTM()

	go rtm.ManageConnection()

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

			case *slack.DisconnectedEvent:
				if ev.Intentional {
					break Loop
				}

			default:
			}
		}
	}
}
