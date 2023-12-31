package configs

import (
	"net/http"

	"github.com/ralvescosta/emqx-sdk-go/credentials"
	"github.com/ralvescosta/emqx-sdk-go/endpoint"
	"go.uber.org/zap"
)

type (
	HTTPClient interface {
		Do(*http.Request) (*http.Response, error)
	}

	Config struct {
		Logger      *zap.SugaredLogger
		Credentials credentials.CredentialsProvider
		Endpoint    endpoint.EndpointResolver
		HTTPClient  HTTPClient
	}
)

func NewConfig() *Config {
	return &Config{}
}

func DefaultConfig() *Config {
	logger, _ := zap.NewProduction()

	return &Config{
		Logger:      logger.Named("emqx-sdk-go").Sugar(),
		Credentials: credentials.NewEnvCredentialsProvider(),
		Endpoint:    endpoint.NewEnvEndpointResolver(),
		HTTPClient:  http.DefaultClient,
	}
}
