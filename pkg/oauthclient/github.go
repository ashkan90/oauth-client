package oauthclient

import (
	"context"
	"encoding/json"
	"oauth-client/pkg/httpclient"
	"oauth-client/pkg/oauthclient/model"

	"github.com/valyala/fasthttp"
)

func (c *client) UserInfoGithub(ctx context.Context, token string) (*model.Generic, error) {
	var resp, err = c.HandleRequest(ctx, httpclient.Request{
		URL: "https://api.github.com/user",
		Headers: map[string]string{
			"Authorization": "token " + token,
		},
		Method: fasthttp.MethodGet,
	})
	if err != nil {
		return nil, err
	}

	var resource model.Generic
	_ = json.Unmarshal(resp.Body, &resource)

	return &resource, err
}
