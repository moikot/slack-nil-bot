# Slack Nil Bot

Nil bot does nothing except keeping WebSocket connection with Slack.

## How to run

You can run it as a Docker container on a Google VM with Docker.

```shell
docker run -d -e SPY_BOT_TOKEN=[token] moikot/slack-nil-bot
```

If you've got Golang environment and Dep, you can build it from source and run.

```shell
git clone git@github.com:moikot/slack-nil-bot.git

cd slack-nil-bot

dep ensure -vendor-only

export export SPY_BOT_TOKEN=[token]; go run .
```
