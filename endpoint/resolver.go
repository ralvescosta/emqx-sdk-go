package endpoint

import "net/url"

type (
	EndpointResolver interface {
		ResolveEndpoint() (*url.URL, error)
	}
)
