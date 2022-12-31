package utils

import "oauth-client/pkg/oauth"

var (
	Providers = map[string]string{
		oauth.ProviderGithub: oauth.GithubCtx,
		oauth.ProviderGoogle: oauth.GoogleCtx,
	}
)
