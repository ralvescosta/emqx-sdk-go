package metrics

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/ralvescosta/emqx-sdk-go/client"
	"github.com/ralvescosta/emqx-sdk-go/service/metrics/types"
	"go.uber.org/zap"
)

func (c *metricsClient) GetMetrics(ctx context.Context, aggregateQueryParam *bool) ([]*types.MetricsResponse, error) {
	urlPath := "/api/v5/metrics"

	query := &url.Values{}
	if aggregateQueryParam != nil {
		query.Add("aggregate", strconv.FormatBool(*aggregateQueryParam))
	} else {
		query = nil
	}

	//urlPath, query, body
	resp, err := c.apiClient.Perform(
		ctx,
		client.NewParams().Get(urlPath).Query(query),
	)
	if err != nil {
		return nil, err
	}

	body := []*types.MetricsResponse{}
	json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		c.logger.Errorw("failure to unmarshal json", zap.Error(err))
		return nil, err
	}

	return body, nil
}
