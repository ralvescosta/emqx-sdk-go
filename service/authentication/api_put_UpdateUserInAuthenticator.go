package authentication

import (
	"context"
	"fmt"
)

func (c *Client) PutUpdateUserInAuthenticator(ctx context.Context, authenticatorID, userID string) error {
	urlPath := fmt.Sprintf("api/v5/authentication/%s/users/%s", authenticatorID, userID)

	if err := c.perform(
		ctx,
		NewRequestParams().Put(urlPath),
	); err != nil {
		return err
	}

	return nil
}
