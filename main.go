package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	e := echo.New()

	e.GET("/callback", handlerMainPage())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/yes", func(c echo.Context) error {
		return c.String(http.StatusOK, "No, World!")
	})
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
	// LINE Botクライアント生成する
	// BOT にはチャネルシークレットとチャネルトークンを環境変数から読み込み引数に渡す

	// res := weather.GetWeather()
	// // テキストメッセージを生成する
	// message := linebot.NewTextMessage(res)
	// // テキストメッセージを友達登録しているユーザー全員に配信する
	// if _, err := bot.BroadcastMessage(message).Do(); err != nil {
	// 	log.Fatal(err)
	// }
}

func handlerMainPage() echo.HandlerFunc {
	return func(c echo.Context) error { //c をいじって Request, Responseを色々する
		bot, err := linebot.New(
			os.Getenv("LINE_BOT_CHANNEL_SECRET"),
			os.Getenv("LINE_BOT_CHANNEL_TOKEN"),
		)
		// エラーに値があればログに出力し終了する
		if err != nil {
			log.Fatal(err)
		}

		events, err := bot.ParseRequest(c.Request())
		if err != nil {
			return nil
		}

		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					replyMessage := message.Text
					if replyMessage == "ぴえん" {
						replyMessage = fmt.Sprintf("ぱおん")
					}
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
						log.Print(err)
					}
				case *linebot.StickerMessage:
					{
						replyMessage := fmt.Sprintf(
							"sticker id is %s, stickerResourceType is %s", message.StickerID, message.StickerResourceType)
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
							log.Print(err)
						}
					}
				}
			}
		}
		return c.String(http.StatusOK, "")
	}
}

func genratestring() string {
	return "ss"
}
