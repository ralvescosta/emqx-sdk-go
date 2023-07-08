package errors

import "errors"

var (
	ErrCredentialsNotProvided  = errors.New("credentials was not provided")
	ErrEnvAPIKeyNotProvided    = errors.New("env EMQX_API_KEY was not provided")
	ErrEnvSecretKeyNotProvided = errors.New("env EMQX_SECRET_KEY was not provided")

	ErrEnvHostNotProvided    = errors.New("env EMQX_HOST was not provided")
	ErrEnvAPIPortNotProvided = errors.New("env EMQX_API_PORT was not provided")
)
