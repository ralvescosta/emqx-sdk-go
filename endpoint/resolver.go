package endpoint

import "net/url"

type (
	Endpoint struct {
		Host string
		Port string
		Path string
		URL  *url.URL
	}

	EndpointResolver interface {
		ResolveEndpoint() (*Endpoint, error)
	}
)
