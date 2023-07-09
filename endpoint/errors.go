package endpoint

import "errors"

var (
	ErrEnvSecretKeyNotProvided = errors.New("env EMQX_SECRET_KEY was not provided")

	ErrEnvHostNotProvided = errors.New("env EMQX_HOST was not provided")
)
