package sso

import (
	"context"
	"oauth-client/pkg/oauth"
	"oauth-client/pkg/oauthclient/model"

	"go.uber.org/zap"
)

type Proxy interface {
	Strategy
}

type proxy struct {
	logger   *zap.Logger
	strategy Strategy
}

func NewStrategyProxy(l *zap.Logger, s Strategy) Proxy {
	return &proxy{
		logger:   l,
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
