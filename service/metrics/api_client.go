package metrics

import (
	"context"

	"github.com/ralvescosta/emqx-sdk-go/client"
	"github.com/ralvescosta/emqx-sdk-go/configs"
	"github.com/ralvescosta/emqx-sdk-go/service/metrics/types"
	"go.uber.org/zap"
)

type (
	MetricsClient interface {
		GetMetrics(ctx context.Context, aggregateQueryParam *bool) ([]*types.MetricsResponse, error)
		GetStats(ctx context.Context, aggregateQueryParam bool) ([]*types.StatsResponse, error)
		GetCurrentStatusGaugeAndRate(ctx context.Context) (*types.MonitorCurrentResponse, error)
		GetNodeCurrentStatusGaugeAndRate(ctx context.Context, nodeParam string) (*types.MonitorCurrentNodeResponse, error)
		GetMonitorDataOnTheNode(ctx context.Context, nodeParam, latestQueryParam string) ([]*types.MonitorNodesResponse, error)
		GetMonitorData(ctx context.Context, latestQueryParam string) ([]*types.MonitorResponse, error)
	}

	metricsClient struct {
		logger    *zap.SugaredLogger
		apiClient *client.APIClient
	}
)

func New(options *client.APIClientOptions) *metricsClient {
	return &metricsClient{options.Logger, client.New(options)}
}

func NewFromConfig(cfg *configs.Config) *metricsClient {
	opts := &client.APIClientOptions{
		Logger:      cfg.Logger,
		Credentials: cfg.Credentials,
		Endpoint:    cfg.Endpoint,
		HTTPClient:  cfg.HTTPClient,
	}

	return New(opts)
}
