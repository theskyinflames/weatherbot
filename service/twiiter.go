package service

import (
	"github.com/ChimeraCoder/anaconda"
)

type (
	Twitter struct {
		log *BotLog
		api *anaconda.TwitterApi
	}
)

func NewTwitter(log *BotLog, cfg *Config) *Twitter {
	return &Twitter{
		log: log,
		api: connectToTwitterApi(cfg),
	}
}

func (t *Twitter) TweetCurrentWeather(weatherStatus string) error {

	tweet, err := t.api.PostTweet(weatherStatus, nil)
	if err != nil {
		return err
	}

	t.log.Debugf("tweeted %s", tweet.FullText)
	return nil
}

func connectToTwitterApi(cfg *Config) *anaconda.TwitterApi {
	anaconda.SetConsumerKey(cfg.TwitterAPICfg.ConsumerKey)
	anaconda.SetConsumerSecret(cfg.TwitterAPICfg.ConsumerSecret)
	return anaconda.NewTwitterApi(cfg.TwitterAPICfg.AccessToken, cfg.TwitterAPICfg.AccessTokenSecret)
}
