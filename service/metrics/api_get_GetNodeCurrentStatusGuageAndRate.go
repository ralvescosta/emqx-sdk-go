package metrics

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/ralvescosta/emqx-sdk-go/client"
	"github.com/ralvescosta/emqx-sdk-go/service/metrics/types"
	"go.uber.org/zap"
)

func (c *metricsClient) GetNodeCurrentStatusGaugeAndRate(ctx context.Context, nodeParam string) (*types.MonitorCurrentNodeResponse, error) {
	if nodeParam == "" {
		return nil, errors.New("node parameter is required")
	}

	urlPath := fmt.Sprintf("/api/v5/monitor_current/nodes/%v", nodeParam)

	resp, err := c.apiClient.Perform(
		ctx,
		client.NewParams().Get(urlPath),
	)
	if err != nil {
		return nil, err
	}

	body := &types.MonitorCurrentNodeResponse{}
	json.NewDecoder(resp.Body).Decode(body)
	if err != nil {
		c.logger.Errorw("failure to unmarshal json", zap.Error(err))

		return nil, err
	}

	return body, nil
}
