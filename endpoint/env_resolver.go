package endpoint

import (
	"fmt"
	"net/url"
	"os"
)

type (
	EnvEndpointResolver struct{}
)

const (
	ENV_HOST     = "EMQX_HOST"
	ENV_API_PORT = "EMQX_API_PORT"
)

func NewEnvEndpointResolver() *EnvEndpointResolver {
	return &EnvEndpointResolver{}
}

func (r EnvEndpointResolver) ResolveEndpoint() (*url.URL, error) {
	host := ""
	port := ""

	if k := os.Getenv(ENV_HOST); k != "" {
		host = k
	} else {
		return nil, ErrEnvHostNotProvided
	}

	if k := os.Getenv(ENV_API_PORT); k != "" {
		port = k
	} else {
		if host[:5] == "http:" {
			port = "80"
		}

		if host[:5] == "https" {
			port = "443"
		}
	}

	u, err := url.Parse(fmt.Sprintf("%v:%v", host, port))
	if err != nil {
		return nil, err
	}

	return u, nil
}
