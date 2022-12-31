package sso

import (
	"context"
	"oauth-client/pkg/oauth"
	"oauth-client/pkg/oauthclient/model"
)

type Strategy interface {
	Login() (string, error)
	Register(context.Context, oauth.CallbackResponse) (*model.Generic, error)
	String() string
}

type StrategySelector interface {
	Algo() Proxy
	AlgoByCtx(c context.Context, provider string) Proxy
	Set(proxy Proxy) Proxy
	SetByCtx(c context.Context, provider string) Proxy
}

type strategy struct {
	algo Proxy
}

func InitSSO(provider Proxy) StrategySelector {
	return &strategy{algo: provider}
}

func (s *strategy) Algo() Proxy {
	return s.algo
}

func (s *strategy) AlgoByCtx(c context.Context, provider string) Proxy {
	return ProxyByContext(c, provider)
}

func (s *strategy) Set(proxy Proxy) Proxy {
	s.algo = proxy
	return s.algo
}

func (s *strategy) SetByCtx(c context.Context, provider string) Proxy {
	s.algo = ProxyByContext(c, provider)
	return s.algo
}
