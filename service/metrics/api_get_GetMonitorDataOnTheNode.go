package metrics

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	"github.com/ralvescosta/emqx-sdk-go/client"
	"github.com/ralvescosta/emqx-sdk-go/service/metrics/types"
	"go.uber.org/zap"
)

func (c *metricsClient) GetMonitorDataOnTheNode(ctx context.Context, nodeParam, latestQueryParam string) ([]*types.MonitorNodesResponse, error) {
	if latestQueryParam == "" {
		return nil, errors.New("latest query parameter is required")
	}

	if nodeParam == "" {
		return nil, errors.New("node parameter is required")
	}

	urlPath := fmt.Sprintf("/api/v5/monitor/nodes/%v", nodeParam)

	query := &url.Values{}
	query.Add("latest", latestQueryParam)

	resp, err := c.apiClient.Perform(
		ctx,
		client.NewParams().Get(urlPath).Query(query),
	)
	if err != nil {
		return nil, err
	}

	body := []*types.MonitorNodesResponse{}
	json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		c.logger.Errorw("failure to unmarshal json", zap.Error(err))
		return nil, err
	}

	return body, nil
}
