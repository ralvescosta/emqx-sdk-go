package metrics

import (
	"context"
	"net/url"
	"strconv"

	"github.com/ralvescosta/emqx-sdk-go/service/metrics/types"
)

func (c *Client) GetStats(ctx context.Context, aggregateQueryParam bool) ([]*types.StatsResponse, error) {
	urlPath := "/api/v5/stats"

	query := &url.Values{}
	query.Add("aggregate", strconv.FormatBool(aggregateQueryParam))

	body := []*types.StatsResponse{}

	err := c.perform(ctx, urlPath, query, body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
