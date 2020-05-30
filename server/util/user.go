package util

import (
	"github.com/k-ueki/tmanager/server/config"
)

func NewUsersClient() *config.TwitterAPIClient {
	conf, token, client := config.Set()

	return &config.TwitterAPIClient{
		Config:     conf,
		Token:      token,
		HttpClient: client,
	}
}
