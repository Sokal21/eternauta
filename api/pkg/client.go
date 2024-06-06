package pkg

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type IHTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	httpClient IHTTPClient
	baseUrl    *url.URL
}

var _ IClient = &Client{}

type httpResponse struct {
	successful bool
	statusCode int
	body       []byte
	headers    http.Header
}

func NewClient(httpClient IHTTPClient, baseUrl string) (*Client, error) {
	parsedUrl, err := url.Parse(baseUrl)
	if err != nil {
		return nil, fmt.Errorf("failed parsing base ulr %s:  %w", baseUrl, err)
	}
	return &Client{
		httpClient: httpClient,
		baseUrl:    parsedUrl,
	}, nil
}

func (c *Client) Get(ctx context.Context, url *url.URL) (*httpResponse, error) {
	req, err := c.prepareRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	return c.do(req)
}

func (c *Client) Post(ctx context.Context, url *url.URL, body []byte) (*httpResponse, error) {

	req, err := c.prepareRequest(ctx, http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}
	return c.do(req)
}

func (c *Client) Put(ctx context.Context, url *url.URL, body []byte) (*httpResponse, error) {
	req, err := c.prepareRequest(ctx, http.MethodPut, url, body)
	if err != nil {
		return nil, err
	}
	return c.do(req)
}

func (c *Client) Delete(ctx context.Context, url *url.URL, body []byte) (*httpResponse, error) {
	req, err := c.prepareRequest(ctx, http.MethodDelete, url, body)
	if err != nil {
		return nil, err
	}
	return c.do(req)
}

func (c *Client) prepareRequest(ctx context.Context, method string, url *url.URL, body []byte) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, url.String(), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	return req, nil
}

func (c *Client) do(req *http.Request) (*httpResponse, error) {
	ctx := req.Context()

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed communicating with Coelsa: %w", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return &httpResponse{
		successful: res.StatusCode >= 200 && res.StatusCode < 300,
		statusCode: res.StatusCode,
		body:       body,
		headers:    res.Header,
	}, nil
}

func (c *Client) ResolvePath(path string) *url.URL {
	return c.baseUrl.ResolveReference(&url.URL{Path: c.baseUrl.Path + path})
}
