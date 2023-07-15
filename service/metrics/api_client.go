package metrics

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/ralvescosta/emqx-sdk-go/configs"
	"github.com/ralvescosta/emqx-sdk-go/credentials"
	"github.com/ralvescosta/emqx-sdk-go/endpoint"
	"go.uber.org/zap"
)

type (
	Options struct {
		Logger      *zap.SugaredLogger
		Credentials credentials.CredentialsProvider
		Endpoint    endpoint.EndpointResolver
		HTTPClient  configs.HTTPClient
	}

	MetricsClient interface{}

	Client struct {
		options *Options
	}
)

func New(options *Options) *Client {
	return &Client{options}
}

func NewFromConfig(cfg *configs.Config) *Client {
	opts := &Options{
		Logger:      cfg.Logger,
		Credentials: cfg.Credentials,
		Endpoint:    cfg.Endpoint,
		HTTPClient:  cfg.HTTPClient,
	}

	return New(opts)
}

func (c *Client) perform(ctx context.Context, path string, queryParams *url.Values, resBody interface{}) error {
	u, err := c.options.Endpoint.ResolveEndpoint()
	if err != nil {
		return err
	}

	u.Path = path
	if queryParams != nil {
		u.RawQuery = queryParams.Encode()
	}

	credentials, err := c.options.Credentials.Retrieve(ctx)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return err
	}

	secHeader, err := credentials.Header()
	if err != nil {
		return err
	}

	req.Header.Set("content-type", "application/json")
	req.Header.Set("authorization", secHeader)

	resp, err := c.options.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&resBody)
	if err != nil {
		return err
	}

	return nil
}
