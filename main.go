package main

import (
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/fizzfuzzHK/line_bot_weather/weather"
	"github.com/line/line-bot-sdk-go/linebot"
)

func handler() {
	res := weather.GetOpenWeather()

	message := linebot.NewTextMessage(res)
	emoji := linebot.NewEmoji(10, "5ac1bfd5040ab15980c9b435", "068")
	message.AddEmoji(emoji)

	bot, err := linebot.New(
		os.Getenv("LINE_BOT_CHANNEL_SECRET"),
		os.Getenv("LINE_BOT_CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := bot.BroadcastMessage(message).Do(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	lambda.Start(handler)
}
