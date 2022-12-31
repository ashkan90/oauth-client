package oauthutil

import (
	"context"
	"oauth-client/pkg/oauth"
	"oauth-client/pkg/sso"
	"oauth-client/pkg/utils"
)

func Google(c context.Context) sso.Proxy {
	return ByProvider(c, oauth.ProviderGoogle)
}

func Github(c context.Context) sso.Proxy {
	return ByProvider(c, oauth.ProviderGithub)
}

func ByProvider(c context.Context, provider string) sso.Proxy {
	return c.Value(utils.Providers[provider]).(sso.Proxy)
}
