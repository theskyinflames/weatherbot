package service

import (
	"context"
	"sync"
	"time"
)

type (
	TweetSender interface {
		TweetCurrentWeather(weatherStatus string) error
	}

	WeatherGetter interface {
		GetWeather(ctx context.Context, city string) (*WeatherStatus, error)
	}

	Service struct {
		log     *BotLog
		twitter TweetSender

		weather WeatherGetter
	}
)

func NewService(log *BotLog, twitter TweetSender, weather WeatherGetter) *Service {
	return &Service{
		log:     log,
		twitter: twitter,
		weather: weather,
	}
}

func (s *Service) Start(ctx context.Context, cfg *Config) error {

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		ticker := time.NewTicker(cfg.ServiceCfg.Interval)
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				s.tweetCurrentWeather(ctx, cfg.WeatherSrvCfg.City)
			}
		}
	}()

	s.log.Info("Service started")
	wg.Wait()
	s.log.Info("Service finished")

	return nil
}

func (s *Service) tweetCurrentWeather(ctx context.Context, city string) {
	currentWeather, err := s.weather.GetWeather(ctx, city)
	if err != nil {
		s.log.Error("Some went wrong when trying to retrive the current weather for the city %s: %s", city, err.Error())
		return
	}

	err = s.twitter.TweetCurrentWeather(currentWeather.String())
	if err != nil {
		s.log.Errorf("Some went wrong when trying to tweet the current weather for the city %s: %s", city, err.Error())
	}

	s.log.Info("Tweet twitted")

	return
}
