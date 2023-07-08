package configs

import (
	"net/http"

	"github.com/ralvescosta/emqx-sdk-go/credentials"
	"go.uber.org/zap"
)

type (
	HTTPClient interface {
		Do(*http.Request) (*http.Response, error)
	}

	Config struct {
		Credentials credentials.CredentialsProvider
		HTTPClient  HTTPClient
		Logger      *zap.SugaredLogger
	}
)

func NewConfig() *Config {
	return &Config{}
}

func DefaultConfig() *Config {
	return &Config{
		Credentials: credentials.NewEnvCredentialsProvider(),
		HTTPClient:  http.DefaultClient,
		Logger:      zap.S().Named("emqx-sdk-go"),
	}
}
