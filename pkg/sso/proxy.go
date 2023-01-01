package sso

import (
	"context"
	"oauth-client/pkg/oauth"
	"oauth-client/pkg/oauthclient/model"
)

type Proxy interface {
	Strategy
}

type proxy struct {
	strategy Strategy
}

func NewStrategyProxy(s Strategy) Proxy {
	return &proxy{
		strategy: s,
	}
}

func (p *proxy) Login() (string, error) {
	return p.strategy.Login()
}

func (p *proxy) Register(c context.Context, r oauth.CallbackResponse) (*model.Generic, error) {
	return p.strategy.Register(c, r)
}

func (p *proxy) String() string {
	return p.strategy.String()
}
