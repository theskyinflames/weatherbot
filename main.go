package main

import (
	"context"
	"net/http"
	"net/url"

	"github.com/theskyinflames/weatherbot/service"
)

func main() {

	log := service.NewBotLog()

	// Load configuration
	cfg := &service.Config{}
	err := cfg.Load()
	if err != nil {
		log.Critical(err.Error())
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	baseurl, err := url.Parse(service.BaseUrl)
	if err != nil {
		log.Critical(err.Error())
	}
	httpClient := service.NewHttpClient(http.DefaultClient, baseurl)
	weather := service.NewWeather(log, httpClient, cfg)

	twitter := service.NewTwitter(log, cfg)

	// Start the service
	srv := service.NewService(log, twitter, weather)
	srv.Start(ctx, cfg)
}
