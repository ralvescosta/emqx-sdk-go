package metrics

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/ralvescosta/emqx-sdk-go/service/metrics/types"
)

func (c *Client) GetMonitorCurrentNode(ctx context.Context, node string) (*types.MonitorCurrentNodeResponse, error) {
	if node == "" {
		return nil, errors.New("node parameter is required")
	}

	url, err := c.options.Endpoint.ResolveEndpoint()
	if err != nil {
		return nil, err
	}

	url.Path = fmt.Sprintf("/api/v5/monitor_current/nodes/%v", node)

	credentials, err := c.options.Credentials.Retrieve(ctx)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	secHeader, err := credentials.Header()
	if err != nil {
		return nil, err
	}

	req.Header.Set("content-type", "application/json")
	req.Header.Set("authorization", secHeader)

	resp, err := c.options.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body := &types.MonitorCurrentNodeResponse{}
	json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func do() {

}
