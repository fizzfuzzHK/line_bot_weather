package main

import (
	"log"
	"net/http"
	"os"

	"github.com/fizzfuzzHK/line_bot_weather/weather"
	"github.com/labstack/echo/v4"
	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/hello", handlerMainPage())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/yes", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
	// LINE Botクライアント生成する
	// BOT にはチャネルシークレットとチャネルトークンを環境変数から読み込み引数に渡す
	bot, err := linebot.New(
		os.Getenv("LINE_BOT_CHANNEL_SECRET"),
		os.Getenv("LINE_BOT_CHANNEL_TOKEN"),
	)
	// エラーに値があればログに出力し終了する
	if err != nil {
		log.Fatal(err)
	}

	res := weather.GetWeather()
	// テキストメッセージを生成する
	message := linebot.NewTextMessage(res)
	// テキストメッセージを友達登録しているユーザー全員に配信する
	if _, err := bot.BroadcastMessage(message).Do(); err != nil {
		log.Fatal(err)
	}
}

func handlerMainPage() echo.handlerFunc {
	return func(c echo.Context) error { //c をいじって Request, Responseを色々する
		str := genratestring()
		return c.String(http.StatusOK, str)
	}
}

func genratestring() string {
	return "ss"
}
