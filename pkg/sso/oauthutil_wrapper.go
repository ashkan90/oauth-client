package sso

import (
	"context"
	"oauth-client/pkg/utils"
)

func ProxyByContext(c context.Context, p string) Proxy {
	return c.Value(utils.Providers[p]).(Proxy)
}
