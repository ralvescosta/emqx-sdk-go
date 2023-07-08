package credentials

import "context"

type (
	CustomCredentialsProvider struct {
		fn func(context.Context) (*Credentials, error)
	}
)

func NewCustomCredentialsProvider(fn func(context.Context) (*Credentials, error)) *CustomCredentialsProvider {
	return &CustomCredentialsProvider{fn}
}

func (p *CustomCredentialsProvider) Retrieve(ctx context.Context) (*Credentials, error) {
	return p.fn(ctx)
}
