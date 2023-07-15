package metrics

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/ralvescosta/emqx-sdk-go/service/metrics/types"
)

func (c *Client) GetMonitorNodes(ctx context.Context, nodeParam, latestQueryParam string) ([]*types.MonitorNodesResponse, error) {
	if latestQueryParam == "" {
		return nil, errors.New("latest query parameter is required")
	}

	if nodeParam == "" {
		return nil, errors.New("node parameter is required")
	}

	body := []*types.MonitorNodesResponse{}

	urlPath := fmt.Sprintf("/api/v5/monitor/nodes/%v", nodeParam)

	query := &url.Values{}
	query.Add("latest", latestQueryParam)

	err := c.perform(ctx, urlPath, query, body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
