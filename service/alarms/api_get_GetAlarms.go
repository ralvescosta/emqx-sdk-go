package alarms

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"

	"github.com/ralvescosta/emqx-sdk-go/client"
	"github.com/ralvescosta/emqx-sdk-go/service/alarms/types"
	"go.uber.org/zap"
)

func (c *alarmsClient) GetAlarms(ctx context.Context, page, limit int, activated *bool) (*types.AlarmsResponse, error) {
	if limit == 0 || page == 0 {
		return nil, errors.New("limit and page must be grater than 0")
	}

	urlPath := "/api/v5/alarms"

	query := &url.Values{}
	query.Add("page", fmt.Sprintf("%v", page))
	query.Add("limit", fmt.Sprintf("%v", limit))

	if activated != nil {
		query.Add("activated", strconv.FormatBool(*activated))
	}

	resp, err := c.apiClient.Perform(
		ctx,
		client.NewParams().Get(urlPath).Query(query),
	)
	if err != nil {
		return nil, err
	}

	body := &types.AlarmsResponse{}
	json.NewDecoder(resp.Body).Decode(body)
	if err != nil {
		c.logger.Errorw("failure to unmarshal json", zap.Error(err))

		return nil, err
	}

	return body, nil
}
