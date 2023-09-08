package client

import (
	"io"
	"net/http"
	"net/url"
)

type (
	Request struct {
		body        io.Reader
		contentType string
	}

	Response struct {
		body any
	}

	Params struct {
		path        string
		method      string
		queryParams *url.Values
		request     *Request
	}
)

func NewParams() *Params {
	return &Params{
		request: &Request{},
	}
}

func (p *Params) Post(path string) *Params {
	p.method = http.MethodPost
	p.path = path

	return p
}

func (p *Params) Get(path string) *Params {
	p.method = http.MethodGet
	p.path = path

	return p
}

func (p *Params) Put(path string) *Params {
	p.method = http.MethodPut
	p.path = path

	return p
}

func (p *Params) Delete(path string) *Params {
	p.method = http.MethodDelete
	p.path = path

	return p
}

func (p *Params) Query(params *url.Values) *Params {
	p.queryParams = params

	return p
}

func (p *Params) RequestBody(body io.Reader) *Params {
	p.request.body = body

	return p
}

func (p *Params) RequestContentType(contentType string) *Params {
	p.request.contentType = contentType

	return p
}
