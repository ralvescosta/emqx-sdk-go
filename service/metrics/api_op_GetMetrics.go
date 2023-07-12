package metrics

import (
	"context"
	"net/http"
)

func (c *Client) GetMetrics(ctx context.Context) error {
	endpoint := c.options.Endpoint.ResolveEndpoint()
	credentials := c.options.Credentials.Retrieve(ctx)

	req, err := http.NewRequest(http.MethodGet, endpoint.URL.String(), nil)
	if err != nil {
		return err
	}

	req.Header.Set("content-type", "application/json")
	req.Header.Set("authorization", credentials.Header())

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	return nil
}
