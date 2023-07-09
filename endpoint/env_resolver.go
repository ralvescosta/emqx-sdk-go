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

func (r EnvEndpointResolver) ResolveEndpoint() (*Endpoint, error) {
	e := &Endpoint{}

	if k := os.Getenv(ENV_HOST); k != "" {
		e.Host = k
	} else {
		return nil, ErrEnvHostNotProvided
	}

	if k := os.Getenv(ENV_API_PORT); k != "" {
		e.Port = k
	} else {
		if e.Host[:5] == "http:" {
			e.Port = "80"
		}

		if e.Host[:5] == "https" {
			e.Port = "443"
		}
	}

	u, err := url.Parse(fmt.Sprintf("%v:%v", e.Host, e.Port))
	if err != nil {
		return nil, err
	}

	e.URL = u
	return e, nil
}
