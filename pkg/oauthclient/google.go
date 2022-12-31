package oauthclient

import (
	"context"
	"encoding/json"
	"oauth-client/pkg/httpclient"
	"oauth-client/pkg/oauthclient/model"

	"github.com/valyala/fasthttp"
)

func (c *client) UserInfoGoogle(ctx context.Context, token string) (*model.Generic, error) {
	var resp, err = c.HandleRequest(ctx, httpclient.Request{
		URL:    prepareURL(token),
		Method: fasthttp.MethodGet,
	})

	if err != nil {
		return nil, err
	}

	var resource model.Generic
	_ = json.Unmarshal(resp.Body, &resource)

	resource.Token = token

	return &resource, nil
}

func prepareURL(token string) string {
	return "https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token
}
