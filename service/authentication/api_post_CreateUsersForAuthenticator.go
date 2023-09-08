package authentication

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/ralvescosta/emqx-sdk-go/client"
	"github.com/ralvescosta/emqx-sdk-go/service/authentication/types"
	"go.uber.org/zap"
)

func (c *authenticationClient) PostCreateUsersForAuthenticator(
	ctx context.Context,
	authenticatorID string,
	req *types.RequestCreateUserForAuthenticator,
) (*types.ResponseCreateUserForAuthenticator, error) {
	urlPath := fmt.Sprintf("api/v5/authentication/%s/users", authenticatorID)

	buffer := new(bytes.Buffer)

	err := json.NewEncoder(buffer).Encode(req)
	if err != nil {
		return nil, err
	}

	resp, err := c.apiClient.Perform(
		ctx,
		client.NewParams().Post(urlPath).RequestBody(buffer),
	)
	if err != nil {
		return nil, err
	}

	body := &types.ResponseCreateUserForAuthenticator{}
	json.NewDecoder(resp.Body).Decode(body)
	if err != nil {
		c.logger.Errorw("failure to unmarshal json", zap.Error(err))
		return nil, err
	}

	return body, nil
}
