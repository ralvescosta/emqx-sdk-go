package authentication

import (
	"context"
	"fmt"

	"github.com/ralvescosta/emqx-sdk-go/client"
)

func (c *authenticationClient) DeleteUserInAuthenticator(ctx context.Context, authenticatorID, userID string) error {
	urlPath := fmt.Sprintf("api/v5/authentication/%s/users/%s", authenticatorID, userID)

	if _, err := c.apiClient.Perform(
		ctx,
		client.NewParams().Delete(urlPath),
	); err != nil {
		return err
	}

	return nil
}
