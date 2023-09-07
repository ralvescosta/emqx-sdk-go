package authentication

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/ralvescosta/emqx-sdk-go/service/authentication/types"
)

func (c *Client) PostCreateUsersForAuthenticator(
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

	res := &types.ResponseCreateUserForAuthenticator{}

	if err := c.perform(
		ctx,
		NewRequestParams().Post(urlPath).ReqBody(buffer).ResBody(res),
	); err != nil {
		return nil, err
	}

	return res, nil
}
