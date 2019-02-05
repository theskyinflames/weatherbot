package service

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const (
	BaseUrl = "http://api.openweathermap.org"
	path    = "/data/2.5/weather"
)

var (
	appid string
)

type (
	HttpClientAccessor interface {
		NewRequest(method, path string, body interface{}, queryValues map[string]string) (*http.Request, error)
		Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error)
	}

	Weather struct {
		log        *BotLog
		httpClient HttpClientAccessor
		appid      string
	}

	// Generated with https://mholt.github.io/json-to-go/ tool
	// From OpenWeatherMap current weather data call https://openweathermap.org/current
	OpenWeatherMapResponse struct {
		Coord struct {
			Lon float64 `json:"lon"`
			Lat float64 `json:"lat"`
		} `json:"coord"`
		Weather []struct {
			ID          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
		Base string `json:"base"`
		Main struct {
			Temp     float64 `json:"temp"`
			Pressure int     `json:"pressure"`
			Humidity int     `json:"humidity"`
			TempMin  float64 `json:"temp_min"`
			TempMax  float64 `json:"temp_max"`
		} `json:"main"`
		Visibility int `json:"visibility"`
		Wind       struct {
			Speed float64 `json:"speed"`
		} `json:"wind"`
		Clouds struct {
			All int `json:"all"`
		} `json:"clouds"`
		Dt  int `json:"dt"`
		Sys struct {
			Type    int     `json:"type"`
			ID      int     `json:"id"`
			Message float64 `json:"message"`
			Country string  `json:"country"`
			Sunrise int     `json:"sunrise"`
			Sunset  int     `json:"sunset"`
		} `json:"sys"`
		ID   int    `json:"id"`
		Name string `json:"name"`
		Cod  int    `json:"cod"`
	}

	WeatherStatus struct {
		Timestamp          time.Time
		City               string
		CurrentTemperature float64
		MinTemperature     float64
		MaxTemperature     float64
		Description        string
	}
)

func NewWeather(log *BotLog, httpClient HttpClientAccessor, cfg *Config) *Weather {
	return &Weather{
		log:        log,
		httpClient: httpClient,
		appid:      cfg.WeatherSrvCfg.AppID,
	}
}

func (ws *WeatherStatus) String() string {
	ts := time.Now()
	return fmt.Sprintf("City: %s, timestamp: %s, temp: %.2f, min temp: %.2f, max temp: %.2f, desc: %s", ws.City, ts.Format(time.RFC3339), ws.CurrentTemperature, ws.MinTemperature, ws.MaxTemperature, ws.Description)
}

func (w *Weather) GetWeather(ctx context.Context, city string) (*WeatherStatus, error) {
	queryValues := map[string]string{
		"q":     city,
		"units": "metric",
		"appid": w.appid,
	}
	fmt.Println("*jas* .... qv ", queryValues)
	req, err := w.httpClient.NewRequest("GET", path, nil, queryValues)
	if err != nil {
		return nil, err
	}

	fmt.Println("*jas* 1")
	var openweathermapResponse OpenWeatherMapResponse
	_, err = w.httpClient.Do(ctx, req, &openweathermapResponse)
	if err != nil {
		return nil, err
	}
	fmt.Println("*jas* 2")
	return &WeatherStatus{
		City:               city,
		CurrentTemperature: openweathermapResponse.Main.Temp,
		MinTemperature:     openweathermapResponse.Main.TempMin,
		MaxTemperature:     openweathermapResponse.Main.TempMin,
		Description:        openweathermapResponse.Weather[0].Description,
	}, err

}
