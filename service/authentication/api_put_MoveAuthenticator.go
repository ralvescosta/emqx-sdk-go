package authentication

import (
	"context"
	"fmt"
)

func (c *Client) PutMoveAuthenticator(ctx context.Context, authenticatorID, position string) error {
	urlPath := fmt.Sprintf("api/v5/authentication/%s/position/%s", authenticatorID, position)

	if err := c.perform(
		ctx,
		NewRequestParams().Put(urlPath),
	); err != nil {

		fmt.Printf("%v", err.Error())
		return err
	}

	return nil
}
