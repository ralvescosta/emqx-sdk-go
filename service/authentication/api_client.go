package authentication

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/ralvescosta/emqx-sdk-go/configs"
	"github.com/ralvescosta/emqx-sdk-go/credentials"
	"github.com/ralvescosta/emqx-sdk-go/endpoint"
	"go.uber.org/zap"
)

type (
	Options struct {
		Logger      *zap.SugaredLogger
		Credentials credentials.CredentialsProvider
		Endpoint    endpoint.EndpointResolver
		HTTPClient  configs.HTTPClient
	}

	AuthenticationClient interface{}

	Client struct {
		options *Options
	}

	RequestParams struct {
		path           string
		method         string
		queryParams    *url.Values
		reqBody        io.Reader
		reqContentType string
		resBody        any
	}
)

func New(options *Options) *Client {
	return &Client{options}
}

func NewFromConfig(cfg *configs.Config) *Client {
	opts := &Options{
		Logger:      cfg.Logger,
		Credentials: cfg.Credentials,
		Endpoint:    cfg.Endpoint,
		HTTPClient:  cfg.HTTPClient,
	}

	return New(opts)
}

func (c *Client) perform(ctx context.Context, params *RequestParams) error {
	u, err := c.options.Endpoint.ResolveEndpoint()
	if err != nil {
		c.options.Logger.Errorw("failure to resolve endpoint", zap.Error(err), zap.String("path", params.path))
		return err
	}

	u.Path = params.path
	if params.queryParams != nil {
		u.RawQuery = params.queryParams.Encode()
	}

	credentials, err := c.options.Credentials.Retrieve(ctx)
	if err != nil {
		c.options.Logger.Errorw("failure to retrieve the api credentials", zap.Error(err))
		return err
	}

	req, err := http.NewRequest(params.method, u.String(), params.reqBody)
	if err != nil {
		c.options.Logger.Errorw("failure to create http request", zap.Error(err), zap.String("path", params.path))
		return err
	}

	secHeader, err := credentials.Header()
	if err != nil {
		c.options.Logger.Errorw("failure to create http headers for authentication", zap.Error(err))
		return err
	}

	req.Header.Set("content-type", params.reqContentType)
	req.Header.Set("authorization", secHeader)

	resp, err := c.options.HTTPClient.Do(req)
	if err != nil {
		c.options.Logger.Errorw("http request error", zap.Error(err), zap.String("path", params.path))
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		c.options.Logger.Errorw("http request error", zap.Int("statusCode", resp.StatusCode))
		return errors.New(resp.Status)
	}

	json.NewDecoder(resp.Body).Decode(&params.resBody)
	if err != nil {
		c.options.Logger.Errorw("failure to unmarshal json", zap.Error(err))
		return err
	}

	return nil
}

func NewRequestParams() *RequestParams {
	return &RequestParams{
		reqContentType: "application/json",
	}
}

func (rp *RequestParams) Post(path string) *RequestParams {
	rp.method = http.MethodPost
	rp.path = path
	return rp
}

func (rp *RequestParams) Get(path string) *RequestParams {
	rp.method = http.MethodGet
	rp.path = path
	return rp
}

func (rp *RequestParams) Put(path string) *RequestParams {
	rp.method = http.MethodPut
	rp.path = path
	return rp
}

func (rp *RequestParams) Delete(path string) *RequestParams {
	rp.method = http.MethodDelete
	rp.path = path
	return rp
}

func (rp *RequestParams) Query(params *url.Values) *RequestParams {
	rp.queryParams = params
	return rp
}

func (rp *RequestParams) ReqBody(body io.Reader) *RequestParams {
	rp.reqBody = body
	return rp
}

func (rp *RequestParams) ReqContentType(contentType string) *RequestParams {
	rp.reqContentType = contentType
	return rp
}

func (rp *RequestParams) ResBody(body any) *RequestParams {
	rp.resBody = body
	return rp
}
