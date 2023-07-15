package metrics

import (
	"context"

	"github.com/ralvescosta/emqx-sdk-go/service/metrics/types"
)

func (c *Client) GetMetrics(ctx context.Context) ([]*types.MetricsResponse, error) {
	urlPath := "/api/v5/metrics"
	body := []*types.MetricsResponse{}

	err := c.perform(ctx, urlPath, nil, body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
