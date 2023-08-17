package authentication

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ralvescosta/emqx-sdk-go/service/authentication/types"
)

func (c *Client) GetUserFromAuthenticator(ctx context.Context, authenticatorID, userID string) (*types.ResponseUserFromAuthenticator, error) {
	urlPath := fmt.Sprintf("api/v5/authentication/%s/users/%s", authenticatorID, userID)

	body := &types.ResponseUserFromAuthenticator{}

	if err := c.perform(ctx, urlPath, http.MethodGet, nil, body); err != nil {
		return nil, err
	}

	return body, nil
}
