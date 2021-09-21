package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Weather struct {
	WeatherType string
	Temp_Max    string
	Temp_Min    string
	Text        string
}

func GetOpenWeather() string {
	url := "http://api.openweathermap.org/data/2.5/onecall"

	API_KEY := os.Getenv("OPEN_WEATHER_API_KEY")
	data := httpRequest(url, API_KEY)
	w := strToJson(data)
	weather := jsonToWeather(w)

	res := makePresentation(weather)

	return res

}

func httpRequest(url string, API_KEY string) string {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	params := req.URL.Query()
	params.Add("lat", "35.652832")
	params.Add("lon", "139.839478")
	params.Add("appid", API_KEY)
	params.Add("exclude", "current,hourly,alerts,minutely")
	req.URL.RawQuery = params.Encode()

	timeout := time.Duration(5 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}

	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(body)
}

func strToJson(data string) map[string]interface{} {
	var weather map[string]interface{}
	if err := json.Unmarshal([]byte(data), &weather); err != nil {
		log.Fatal("JSON Unmarshall Error", err)
	}
	return weather
}

func jsonToWeather(w map[string]interface{}) *Weather {
	weather := new(Weather)
	icon := w["daily"].([]interface{})[0].(map[string]interface{})["weather"].([]interface{})[0].(map[string]interface{})["icon"].(string)
	weather.WeatherType = getIcon(icon)
	temp_max := w["daily"].([]interface{})[0].(map[string]interface{})["temp"].(map[string]interface{})["max"].(float64)
	temp_min := w["daily"].([]interface{})[0].(map[string]interface{})["temp"].(map[string]interface{})["min"].(float64)

	weather.Temp_Max = kelvinToCelsius(temp_max)
	weather.Temp_Min = kelvinToCelsius(temp_min)

	return weather
}

func kelvinToCelsius(Kelvin float64) string {
	celsius := math.Round((Kelvin-273.15)*10) / 10
	strCelsius := strconv.FormatFloat(celsius, 'f', 1, 64)
	return strCelsius
}

func makePresentation(w *Weather) string {
	weatherType := fmt.Sprintf("おはようございます！$\n本日の東京の天気は%sです\n", w.WeatherType)
	max := fmt.Sprintf("最高気温：%s度\n", w.Temp_Max)
	min := fmt.Sprintf("最低気温：%s度", w.Temp_Min)
	return string(weatherType + max + min)
}

func getIcon(icon string) string {
	switch icon {
	case "01d", "02d":
		return "☀"
	case "03d", "04d":
		return "☁"
	case "9d", "10d":
		return "🌧"
	case "11d":
		return "🌩"
	case "13d":
		return "🌨"
	case "50d":
		return "🌫"
	default:
		return "？"
	}
}
