package metrics

import (
	"github.com/ralvescosta/emqx-sdk-go/credentials"
)

type (
	Options struct {
		Credentials *credentials.CredentialsProvider
	}

	MetricsClient interface{}
	Client        struct{}
)

func New(options *Options) *Client {
	return nil
}

func NewFromConfig() *Client {
	return nil
}
