package credentials

import (
	"context"
	"os"

	"github.com/ralvescosta/emqx-sdk-go/errors"
)

type (
	EnvCredentialsProvider struct{}
)

const (
	ENV_API_KEY    = "EMQX_API_KEY"
	ENV_SECRET_KEY = "EMQX_SECRET_KEY"
)

func NewEnvCredentialsProvider() *EnvCredentialsProvider {
	return &EnvCredentialsProvider{}
}

func (p *EnvCredentialsProvider) Retrieve(ctx context.Context) (*Credentials, error) {
	c := &Credentials{}

	if k := os.Getenv(ENV_API_KEY); k != "" {
		c.APIKey = k
	} else {
		return nil, errors.ErrEnvAPIKeyNotProvided
	}

	if k := os.Getenv(ENV_SECRET_KEY); k != "" {
		c.SecretKey = k
	} else {
		return nil, errors.ErrEnvSecretKeyNotProvided
	}

	return c, nil
}
