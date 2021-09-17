package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Weather struct {
	targetArea   string
	headlineText string
	text         string
}

func GetWeather(w *Weather) string {
	url := "https://www.jma.go.jp/bosai/forecast/data/overview_forecast/130000.json"
	data := httpRequest(url)
	weather := strToJson(data)
	area := fmt.Sprintf("今日の天気は %s", weather.area)
	head := fmt.Sprintf("今日の天気は %s", weather.headlineText)
	text := fmt.Sprintf("今日の天気は %s", weather.text)
	response := area + head + text
	return response
}

func httpRequest(url string) string {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal("Get Http Error:", err)
	}
	// レスポンスボディを読み込む
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal("IO Read Error:", err)
	}
	// 読み込み終わったらレスポンスボディを閉じる
	defer response.Body.Close()
	return string(body)
}

func strToJson(data string) *Weather {
	weather := new(Weather)
	if err := json.Unmarshal([]byte(data), weather); err != nil {
		log.Fatal("JSON Unmarshall Error", err)
	}
	return weather
}
