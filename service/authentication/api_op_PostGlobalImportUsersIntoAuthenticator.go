package authentication

import (
	"context"
	"fmt"
	"net/http"
)

// Multi part form data
func (c *Client) PostGlobalImportUsersIntoAuthenticator(ctx context.Context, authenticatorID string) error {
	urlPath := fmt.Sprintf("api/v5/authentication/%s/import_users", authenticatorID)

	if err := c.perform(ctx, urlPath, http.MethodPost, nil, nil); err != nil {
		return err
	}

	return nil
}
