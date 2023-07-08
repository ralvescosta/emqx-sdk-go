package endpoint

import (
	"os"

	"github.com/ralvescosta/emqx-sdk-go/errors"
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

func (r EnvEndpointResolver) ResolveEndpoint(service string) (*Endpoint, error) {
	e := &Endpoint{}

	if k := os.Getenv(ENV_HOST); k != "" {
		e.Host = k
	} else {
		return nil, errors.ErrEnvHostNotProvided
	}

	//use uri package to ensure the host was provided correctly and to use to add the other parameters

	if k := os.Getenv(ENV_API_PORT); k != "" {
		e.Port = k
	} else {
		return nil, errors.ErrEnvAPIPortNotProvided
	}

	return e, nil
}
