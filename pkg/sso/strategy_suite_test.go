package sso_test

import (
	"context"
	om "oauth-client/pkg/oauth/mocks"
	"oauth-client/pkg/oauthclient/model"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"golang.org/x/oauth2"
)

func TestSSO(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SSO Unit Tests")
}

const (
	githubInvalidLoginURL = "https://github.com/login/oauth/authorize?client_id=invalid-client-id&redirect_uri=http%3A%2F%2Flocalhost%2Fregister%3Fprovider%3Dgithub&response_type=code&scope=&state=stateString"
	googleInvalidLoginURL = "https://accounts.google.com/o/oauth2/auth?client_id=&redirect_uri=http%3A%2F%2Flocalhost%2Fregister%3Fprovider%3Dgoogle&response_type=code&scope=&state=stateString"
)

var (
	ctx             context.Context
	oauthTok        *oauth2.Token
	mockCtrl        *gomock.Controller
	googleProcessor *om.MockProcessor
	githubProcessor *om.MockProcessor
)

var _ = BeforeEach(func() {
	ctx = context.Background()
	mockCtrl = gomock.NewController(GinkgoT())

	oauthTok = &oauth2.Token{
		AccessToken:  "token",
		TokenType:    "type",
		RefreshToken: "refresh-token",
		Expiry:       time.Now().Add(time.Minute * 1),
	}
	googleProcessor = om.NewMockProcessor(mockCtrl)
	githubProcessor = om.NewMockProcessor(mockCtrl)
})

func successByProvider(provider *om.MockProcessor, ex *model.Generic) {
	provider.EXPECT().GetExchange(ctx, gomock.Any()).Return(oauthTok, nil).AnyTimes()
	provider.EXPECT().GetUserInfo(ctx, oauthTok).Return(ex, nil).AnyTimes()
}
