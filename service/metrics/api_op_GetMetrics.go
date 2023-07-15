package metrics

import (
	"context"
	"net/url"
	"strconv"

	"github.com/ralvescosta/emqx-sdk-go/service/metrics/types"
)

func (c *Client) GetMetrics(ctx context.Context, aggregateQueryParam *bool) ([]*types.MetricsResponse, error) {
	urlPath := "/api/v5/metrics"

	query := &url.Values{}
	if aggregateQueryParam != nil {
		query.Add("aggregate", strconv.FormatBool(*aggregateQueryParam))
	} else {
		query = nil
	}

	body := []*types.MetricsResponse{}

	err := c.perform(ctx, urlPath, query, body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
