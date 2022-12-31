package oauthclient

import (
	"context"
	"oauth-client/pkg/httpclient"
	"oauth-client/pkg/oauthclient/model"
)

type Client interface {
	HandleRequest(ctx context.Context, req httpclient.Request) (*httpclient.Response, error)

	UserInfoGoogle(ctx context.Context, token string) (*model.Generic, error)
	UserInfoGithub(ctx context.Context, token string) (*model.Generic, error)
}

type client struct {
	httpClient httpclient.HTTPClient
}

func NewClient(hc httpclient.HTTPClient) Client {
	return &client{
		httpClient: hc,
	}
}

func (c *client) HandleRequest(ctx context.Context, req httpclient.Request) (*httpclient.Response, error) {
	resp, err := c.httpClient.HandleRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	if c.httpClient.IsSuccessStatusCode(resp) {
		return resp, nil
	}

	return nil, c.httpClient.HandleException(resp)
}
