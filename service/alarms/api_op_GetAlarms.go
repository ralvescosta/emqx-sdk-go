package alarms

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"strconv"

	"github.com/ralvescosta/emqx-sdk-go/service/alarms/types"
)

func (c *Client) GetAlarms(ctx context.Context, page, limit int, activated *bool) ([]*types.AlarmsResponse, error) {
	if limit == 0 {
		return nil, errors.New("limit must be grater than 0")
	}

	urlPath := "/api/v5/metrics"
	query := &url.Values{}

	query.Add("page", fmt.Sprintf("%v", page))
	query.Add("limit", fmt.Sprintf("%v", limit))

	if activated != nil {
		query.Add("activated", strconv.FormatBool(*activated))
	}

	body := []*types.AlarmsResponse{}

	err := c.perform(ctx, urlPath, query, body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
