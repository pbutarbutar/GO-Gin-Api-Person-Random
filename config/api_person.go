package config

import "os"

type ApiPersonConfig struct {
	Url string
}

func NewApiPersonConfig() *ApiPersonConfig {
	apiUrl := os.Getenv("API_PERSON_URL")
	if apiUrl == "" {
		apiUrl = "https://randomuser.me/api/"
	}
	return &ApiPersonConfig{
		Url: apiUrl,
	}
}
