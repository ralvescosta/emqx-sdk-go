package authentication

import (
	"context"
	"fmt"

	"github.com/ralvescosta/emqx-sdk-go/client"
)

func (c *authenticationClient) PutMoveAuthenticator(ctx context.Context, authenticatorID, position string) error {
	urlPath := fmt.Sprintf("api/v5/authentication/%s/position/%s", authenticatorID, position)

	_, err := c.apiClient.Perform(
		ctx,
		client.NewParams().Put(urlPath),
	)

	if err != nil {
		fmt.Printf("%v", err.Error())

		return err
	}

	return nil
}
