package metrics

import (
	"context"
	"encoding/json"
	"errors"
	"net/url"

	"github.com/ralvescosta/emqx-sdk-go/client"
	"github.com/ralvescosta/emqx-sdk-go/service/metrics/types"
	"go.uber.org/zap"
)

func (c *metricsClient) GetListMonitorData(ctx context.Context, latestQueryParam string) ([]*types.MonitorResponse, error) {
	if latestQueryParam == "" {
		return nil, errors.New("latest query parameter is required")
	}

	urlPath := "/api/v5/monitor"

	query := &url.Values{}
	query.Add("latest", latestQueryParam)

	resp, err := c.apiClient.Perform(
		ctx,
		client.NewParams().Get(urlPath).Query(query),
	)
	if err != nil {
		return nil, err
	}

	body := []*types.MonitorResponse{}
	json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		c.logger.Errorw("failure to unmarshal json", zap.Error(err))

		return nil, err
	}

	return body, nil
}
