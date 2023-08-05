package metrics

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/ralvescosta/emqx-sdk-go/configs"
	"github.com/ralvescosta/emqx-sdk-go/credentials"
	"github.com/ralvescosta/emqx-sdk-go/endpoint"
	"github.com/ralvescosta/emqx-sdk-go/service/metrics/types"
	"go.uber.org/zap"
)

type (
	Options struct {
		Logger      *zap.SugaredLogger
		Credentials credentials.CredentialsProvider
		Endpoint    endpoint.EndpointResolver
		HTTPClient  configs.HTTPClient
	}

	MetricsClient interface {
		GetMetrics(ctx context.Context, aggregateQueryParam *bool) ([]*types.MetricsResponse, error)
		GetMonitor(ctx context.Context, latestQueryParam string) ([]*types.MonitorResponse, error)
		GetMonitorCurrent(ctx context.Context) (*types.MonitorCurrentResponse, error)
		GetMonitorCurrentNode(ctx context.Context, nodeParam string) (*types.MonitorCurrentNodeResponse, error)
		GetMonitorNodes(ctx context.Context, nodeParam, latestQueryParam string) ([]*types.MonitorNodesResponse, error)
		GetStats(ctx context.Context, aggregateQueryParam bool) ([]*types.StatsResponse, error)
	}

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
		c.options.Logger.Errorw("failure to resolve endpoint", zap.Error(err), zap.String("path", path))
		return err
	}

	u.Path = path
	if queryParams != nil {
		u.RawQuery = queryParams.Encode()
	}

	credentials, err := c.options.Credentials.Retrieve(ctx)
	if err != nil {
		c.options.Logger.Errorw("failure to retrieve the api credentials", zap.Error(err))
		return err
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		c.options.Logger.Errorw("failure to create http request", zap.Error(err), zap.String("path", path))
		return err
	}

	secHeader, err := credentials.Header()
	if err != nil {
		c.options.Logger.Errorw("failure to create http headers for authentication", zap.Error(err))
		return err
	}

	req.Header.Set("content-type", "application/json")
	req.Header.Set("authorization", secHeader)

	resp, err := c.options.HTTPClient.Do(req)
	if err != nil {
		c.options.Logger.Errorw("http request error", zap.Error(err), zap.String("path", path))
		return err
	}
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&resBody)
	if err != nil {
		c.options.Logger.Errorw("failure to unmarshal json", zap.Error(err))
		return err
	}

	return nil
}
