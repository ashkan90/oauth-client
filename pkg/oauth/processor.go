package oauth

import (
	"context"
	neturl "net/url"
	oauth_client "oauth-client"
	"oauth-client/pkg/oauthclient"
	"oauth-client/pkg/oauthclient/model"
	"strings"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
)

type Processor interface {
	BuildLoginURL() (string, error)
	GetExchange(c context.Context, resp CallbackResponse) (*oauth2.Token, error)
	GetUserInfo(c context.Context, tok *oauth2.Token) (*model.Generic, error)
}

type processor struct {
	client oauthclient.Client
	conf   *oauth2.Config
}

func NewOAuthProcessor(client oauthclient.Client, conf oauth_client.GeneralOAuthSettings) Processor {
	var endpoint = google.Endpoint
	if strings.HasSuffix(conf.RedirectURL, "github") {
		endpoint = github.Endpoint
	}

	return &processor{
		client: client,
		conf: &oauth2.Config{
			ClientID:     conf.ClientID,
			ClientSecret: conf.ClientSecret,
			RedirectURL:  conf.RedirectURL,
			Scopes:       conf.Scopes,
			Endpoint:     endpoint,
		},
	}
}

func (p processor) BuildLoginURL() (string, error) {
	var (
		url, err = neturl.Parse(p.conf.Endpoint.AuthURL)
		params   = neturl.Values{
			LoginCredentialClientID:     []string{p.conf.ClientID},
			LoginCredentialRedirectURI:  []string{p.conf.RedirectURL},
			LoginCredentialResponseType: []string{"code"},
			LoginCredentialState:        []string{"stateString"},
			LoginCredentialScope:        []string{strings.Join(p.conf.Scopes, " ")},
		}
	)

	if err != nil {
		return "", err
	}

	url.RawQuery = params.Encode()

	return url.String(), nil
}

func (p processor) GetExchange(c context.Context, resp CallbackResponse) (*oauth2.Token, error) {
	var tok, err = p.conf.Exchange(c, resp.Code)
	if err != nil {
		return nil, err
	}

	return tok, nil
}

func (p processor) GetUserInfo(c context.Context, tok *oauth2.Token) (*model.Generic, error) {
	var (
		res *model.Generic
		err error
	)

	if strings.HasSuffix(p.conf.RedirectURL, "github") {
		res, err = p.client.UserInfoGithub(c, tok.AccessToken)
	} else {
		res, err = p.client.UserInfoGoogle(c, tok.AccessToken)
	}

	if err != nil {
		return nil, err
	}

	return res, nil
}
