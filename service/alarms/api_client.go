package alarms

import (
	"context"

	"github.com/ralvescosta/emqx-sdk-go/client"
	"github.com/ralvescosta/emqx-sdk-go/configs"
	"github.com/ralvescosta/emqx-sdk-go/service/alarms/types"
	"go.uber.org/zap"
)

type (
	AlarmsClient interface {
		DeleteRemoveAllHistoricalAlarms(ctx context.Context) error
		GetListAlarms(ctx context.Context, page, limit int, activated *bool) (*types.AlarmsResponse, error)
	}

	alarmsClient struct {
		logger    *zap.SugaredLogger
		apiClient *client.APIClient
	}
)

func New(options *client.APIClientOptions) *alarmsClient {
	return &alarmsClient{options.Logger, client.New(options)}
}

func NewFromConfig(cfg *configs.Config) *alarmsClient {
	opts := &client.APIClientOptions{
		Logger:      cfg.Logger,
		Credentials: cfg.Credentials,
		Endpoint:    cfg.Endpoint,
		HTTPClient:  cfg.HTTPClient,
	}

	return New(opts)
}
