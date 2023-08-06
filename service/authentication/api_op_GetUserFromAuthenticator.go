package authentication

import (
	"context"
	"fmt"
	"net/http"
)

func (c *Client) GetUserFromAuthenticator(ctx context.Context, authenticatorID, userID string) error {
	urlPath := fmt.Sprintf("api/v5/authentication/%s/users/%s", authenticatorID, userID)

	if err := c.perform(ctx, urlPath, http.MethodGet, nil, nil); err != nil {
		return err
	}

	return nil
}
