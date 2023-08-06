package authentication

import (
	"context"
	"fmt"
	"net/http"
)

func (c *Client) DeleteAuthenticationUser(ctx context.Context, authenticatorID, userID string) error {
	urlPath := fmt.Sprintf("api/v5/authentication/%s/users/%s", authenticatorID, userID)

	if err := c.perform(ctx, urlPath, http.MethodDelete, nil, nil); err != nil {
		return err
	}

	return nil
}
