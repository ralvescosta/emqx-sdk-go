package metrics

import (
	"context"
	"errors"
	"net/url"

	"github.com/ralvescosta/emqx-sdk-go/service/metrics/types"
)

func (c *Client) GetMonitor(ctx context.Context, latestQueryParam string) ([]*types.MonitorResponse, error) {
	if latestQueryParam == "" {
		return nil, errors.New("latest query parameter is required")
	}

	body := []*types.MonitorResponse{}

	urlPath := "/api/v5/monitor"

	query := &url.Values{}
	query.Add("latest", latestQueryParam)

	err := c.perform(ctx, urlPath, query, body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
