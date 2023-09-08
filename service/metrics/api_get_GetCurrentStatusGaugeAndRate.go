package metrics

import (
	"context"
	"encoding/json"

	"github.com/ralvescosta/emqx-sdk-go/client"
	"github.com/ralvescosta/emqx-sdk-go/service/metrics/types"
	"go.uber.org/zap"
)

func (c *metricsClient) GetCurrentStatusGaugeAndRate(ctx context.Context) (*types.MonitorCurrentResponse, error) {
	urlPath := "/api/v5/monitor_current"

	resp, err := c.apiClient.Perform(ctx, client.NewParams().Get(urlPath))
	if err != nil {
		return nil, err
	}

	body := &types.MonitorCurrentResponse{}
	json.NewDecoder(resp.Body).Decode(body)
	if err != nil {
		c.logger.Errorw("failure to unmarshal json", zap.Error(err))
		return nil, err
	}

	return body, nil
}
