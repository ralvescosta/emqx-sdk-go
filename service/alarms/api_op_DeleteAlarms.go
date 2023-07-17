package alarms

import (
	"context"
	"net/http"
)

func (c *Client) DeleteAlarms(ctx context.Context) error {
	urlPath := "/api/v5/alarms"

	body := ""

	err := c.perform(ctx, urlPath, http.MethodDelete, nil, &body)
	if err != nil {
		return err
	}

	return nil
}
