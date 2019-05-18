# Slack Nil Bot
[![Build Status](https://travis-ci.com/moikot/slack-nil-bot.svg?branch=master)](https://travis-ci.com/moikot/slack-nil-bot)
[![Go Report Card](https://goreportcard.com/badge/github.com/moikot/slack-nil-bot)](https://goreportcard.com/report/github.com/moikot/slack-nil-bot)

Nil bot does nothing except keeping WebSocket connection with Slack.

## How to run

### Using Kubernetes and Helm

If you are running Kubernetes with Helm installed you can start nil-bot in this way:

```shell
helm repo add moikot https://moikot.github.io/helm-charts
helm install moikot/slack-nil-bot --name=nil-bot --set token=[bot user key]

```


### Using Docker

On Raspberry Pi (arm) or Linux (amd64) or you can run it on Docker.

```shell
docker run -d --restart unless-stopped -e BOT_TOKEN=[bot user token] moikot/slack-nil-bot

``` 

### Compile and run from code

If you've got Golang environment and Dep installed, you can build it from source and run.

```shell
git clone git@github.com:moikot/slack-nil-bot.git

cd slack-nil-bot

dep ensure -vendor-only

export export BOT_TOKEN=[token]; go run .
```

## Prometheus

Nil Bot exposes Prometheus counters on `0.0.0.0:9153/metrics` endpoint.
If you have Prometheus and Grafana running in your Kubernetes cluster you can add 
`nil_bot_connection_status` gauge on a dashboard. The value of gauge is 
one when Nil Bot is connected to Slack, and zero otherwise.