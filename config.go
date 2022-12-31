package oauth_client

type AuthConfig struct {
	Google GeneralOAuthSettings `mapstructure:"google"`
	Github GeneralOAuthSettings `mapstructure:"github"`
}

type GeneralOAuthSettings struct {
	Scopes       []string `mapstructure:"scopes"`
	RedirectURL  string   `mapstructure:"redirect-url"`
	ClientID     string   `mapstructure:"client-id"`
	ClientSecret string   `mapstructure:"client-secret"`
}
