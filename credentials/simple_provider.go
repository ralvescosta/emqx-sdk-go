package credentials

import (
	"context"
)

type (
	SimpleCredentialsProvider struct {
		apiKey    string
		secretKey string
	}
)

func NewSimpleCredentialsProvider(apiKey, secretKey string) *SimpleCredentialsProvider {
	return &SimpleCredentialsProvider{}
}

func (p *SimpleCredentialsProvider) Retrieve(ctx context.Context) (*Credentials, error) {
	return &Credentials{
		APIKey:    p.apiKey,
		SecretKey: p.secretKey,
	}, nil
}
