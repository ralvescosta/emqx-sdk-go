package authentication

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ralvescosta/emqx-sdk-go/client"
	"github.com/ralvescosta/emqx-sdk-go/service/authentication/types"
	"go.uber.org/zap"
)

func (c *authenticationClient) GetUserFromAuthenticator(ctx context.Context, authenticatorID, userID string) (*types.ResponseUserFromAuthenticator, error) {
	urlPath := fmt.Sprintf("api/v5/authentication/%s/users/%s", authenticatorID, userID)

	resp, err := c.apiClient.Perform(
		ctx,
		client.NewParams().Get(urlPath),
	)
	if err != nil {
		return nil, err
	}

	body := &types.ResponseUserFromAuthenticator{}
	json.NewDecoder(resp.Body).Decode(body)
	if err != nil {
		c.logger.Errorw("failure to unmarshal json", zap.Error(err))

		return nil, err
	}

	return body, nil
}
