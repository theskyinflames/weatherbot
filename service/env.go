package service

import (
	"net/url"
	"os"
	"strconv"
	"time"
)

type (
	ServiceCfg struct {
		Interval time.Duration
	}

	WeatherSrvCfg struct {
		AppID string
		City  string
	}

	TwitterAPICfg struct {
		ConsumerKey       string
		ConsumerSecret    string
		AccessToken       string
		AccessTokenSecret string
	}

	Config struct {
		ServiceCfg
		WeatherSrvCfg
		TwitterAPICfg
	}
)

func (c *Config) Load() error {

	interval, err := strconv.Atoi(getEnv("SERVICE_PUBLISHING_INTERVAL"))
	if err != nil {
		return err
	}
	c.ServiceCfg.Interval = time.Duration(interval) * time.Second

	c.WeatherSrvCfg.City = url.QueryEscape(getEnv("WEATHER_CITY"))
	c.WeatherSrvCfg.AppID = getEnv("OPENWEATHERMAP_APPID")

	c.TwitterAPICfg.ConsumerKey = getEnv("TWITTER_CONSUMER_KEY")
	c.TwitterAPICfg.ConsumerSecret = getEnv("TWITTER_CONSUMER_SECRET")
	c.TwitterAPICfg.AccessToken = getEnv("TWITTER_ACCESS_TOKEN")
	c.TwitterAPICfg.AccessTokenSecret = getEnv("TWITTER_ACCESS_TOKEN_SECRET")

	return nil
}

func getEnv(env string) (value string) {
	value = os.Getenv(env)
	if len(value) == 0 {
		panic("environment variable " + env + " does not exist")
	}
	return
}
