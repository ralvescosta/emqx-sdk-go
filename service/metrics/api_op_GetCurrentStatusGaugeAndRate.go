package metrics

import (
	"context"

	"github.com/ralvescosta/emqx-sdk-go/service/metrics/types"
)

func (c *Client) GetCurrentStatusGaugeAndRate(ctx context.Context) (*types.MonitorCurrentResponse, error) {
	urlPath := "/api/v5/monitor_current"
	body := &types.MonitorCurrentResponse{}

	err := c.perform(ctx, urlPath, nil, body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
