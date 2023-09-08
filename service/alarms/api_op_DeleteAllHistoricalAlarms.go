package alarms

import (
	"context"

	"github.com/ralvescosta/emqx-sdk-go/client"
)

func (c *alarmsClient) DeleteAllHistoricalAlarms(ctx context.Context) error {
	urlPath := "/api/v5/alarms"

	_, err := c.apiClient.Perform(
		ctx,
		client.NewParams().Delete(urlPath),
	)
	if err != nil {
		return err
	}

	return nil
}
