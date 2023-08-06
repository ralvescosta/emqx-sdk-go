package authentication

import (
	"context"
	"fmt"
	"net/http"
)

func (c *Client) PutMoveAuthenticator(ctx context.Context, authenticatorID, position string) error {
	urlPath := fmt.Sprintf("api/v5/authentication/%s/position/%s", authenticatorID, position)

	if err := c.perform(ctx, urlPath, http.MethodPut, nil, nil); err != nil {

		fmt.Printf("%v", err.Error())
		return err
	}

	return nil
}
