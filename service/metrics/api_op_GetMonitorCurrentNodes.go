package metrics

import (
	"context"
	"errors"
	"fmt"

	"github.com/ralvescosta/emqx-sdk-go/service/metrics/types"
)

func (c *Client) GetMonitorCurrentNode(ctx context.Context, nodeParam string) (*types.MonitorCurrentNodeResponse, error) {
	if nodeParam == "" {
		return nil, errors.New("node parameter is required")
	}

	urlPath := fmt.Sprintf("/api/v5/monitor_current/nodes/%v", nodeParam)
	body := &types.MonitorCurrentNodeResponse{}

	err := c.perform(ctx, urlPath, nil, body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
