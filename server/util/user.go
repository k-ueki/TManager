package util

import (
	"github.com/k-ueki/tmanager/server/config"
)

func NewUsersClient() *config.Client {
	conf, token, client := config.Set()

	return &config.Client{
		Config:     conf,
		Token:      token,
		HttpClient: client,
	}
}
