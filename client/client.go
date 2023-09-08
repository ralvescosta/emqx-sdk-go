package client

import (
	"context"
	"net/http"

	"github.com/ralvescosta/emqx-sdk-go/configs"
	"github.com/ralvescosta/emqx-sdk-go/credentials"
	"github.com/ralvescosta/emqx-sdk-go/endpoint"
	"go.uber.org/zap"
)

type (
	APIClientOptions struct {
		Logger      *zap.SugaredLogger
		Credentials credentials.CredentialsProvider
		Endpoint    endpoint.EndpointResolver
		HTTPClient  configs.HTTPClient
	}

	APIClient struct {
		opts *APIClientOptions
	}
)

func New(options *APIClientOptions) *APIClient {
	return &APIClient{options}
}

func NewFromConfig(cfg *configs.Config) *APIClient {
	opts := &APIClientOptions{
		Logger:      cfg.Logger,
		Credentials: cfg.Credentials,
		Endpoint:    cfg.Endpoint,
		HTTPClient:  cfg.HTTPClient,
	}

	return New(opts)
}

func (c *APIClient) Perform(ctx context.Context, params *Params) (*http.Response, error) {
	u, err := c.opts.Endpoint.ResolveEndpoint()
	if err != nil {
		c.opts.Logger.Errorw("failure to resolve endpoint", zap.Error(err), zap.String("path", params.path))

		return nil, err
	}

	u.Path = params.path
	if params.queryParams != nil {
		u.RawQuery = params.queryParams.Encode()
	}

	credentials, err := c.opts.Credentials.Retrieve(ctx)
	if err != nil {
		c.opts.Logger.Errorw("failure to retrieve the api credentials", zap.Error(err))

		return nil, err
	}

	req, err := http.NewRequest(params.method, u.String(), params.request.body)
	if err != nil {
		c.opts.Logger.Errorw("failure to create http request", zap.Error(err), zap.String("path", params.path))

		return nil, err
	}

	secHeader, err := credentials.Header()
	if err != nil {
		c.opts.Logger.Errorw("failure to create http headers for authentication", zap.Error(err))

		return nil, err
	}

	req.Header.Set("content-type", params.request.contentType)
	req.Header.Set("authorization", secHeader)

	resp, err := c.opts.HTTPClient.Do(req)
	if err != nil {
		c.opts.Logger.Errorw("http request error", zap.Error(err), zap.String("path", params.path))

		return nil, err
	}

	if resp.StatusCode >= 400 {
		c.opts.Logger.Errorw("http request error", zap.Int("statusCode", resp.StatusCode))
		err := NewErrorFromStatusCode(resp.StatusCode)

		return nil, err
	}

	return resp, nil
}
