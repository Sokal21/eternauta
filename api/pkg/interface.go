package pkg

import (
	"context"
	"net/url"
)

type IClient interface {
	Get(ctx context.Context, url *url.URL) (*httpResponse, error)
	Post(ctx context.Context, url *url.URL, body []byte) (*httpResponse, error)
	Put(ctx context.Context, url *url.URL, body []byte) (*httpResponse, error)
	Delete(ctx context.Context, url *url.URL, body []byte) (*httpResponse, error)
	ResolvePath(path string) *url.URL
}
