package endpoint

type (
	Endpoint struct {
		Host string
		Port string
		Path string
		URL  string
	}

	EndpointResolver interface {
		ResolveEndpoint(service string) (*Endpoint, error)
	}
)
