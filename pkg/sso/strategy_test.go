package sso_test

import (
	"oauth-client/pkg/oauth"
	"oauth-client/pkg/oauthclient/model"
	"oauth-client/pkg/sso"
	"oauth-client/pkg/sso/strategies"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("SSO Unit Tests", func() {
	Describe("init sso with given correct provider", func() {
		It("should return new github sso", func() {
			_sso := sso.InitSSO(sso.NewStrategyProxy(
				strategies.NewGithubSSO(githubProcessor),
			))

			Expect(_sso).To(Not(BeNil()))
		})

		It("should return new google sso", func() {
			_sso := sso.InitSSO(sso.NewStrategyProxy(
				strategies.NewGoogleSSO(googleProcessor),
			))

			Expect(_sso).To(Not(BeNil()))
		})
	})

	Describe("initiated sso provider can be changed at runtime", func() {
		var _sso sso.StrategySelector
		BeforeEach(func() {
			_sso = sso.InitSSO(sso.NewStrategyProxy(
				strategies.NewGithubSSO(githubProcessor),
			))
		})

		It("should be github sso by default", func() {
			Expect(_sso.Algo()).To(Not(BeNil()))
			Expect(_sso.Algo().String()).To(Equal("github-sso"))
		})

		It("should set current sso provider by given provider", func() {
			googleProvider := sso.NewStrategyProxy(strategies.NewGoogleSSO(googleProcessor))
			_sso.Set(googleProvider)

			Expect(_sso.Algo()).To(Not(BeNil()))
			Expect(_sso.Algo().String()).To(Equal("google-sso"))
		})
	})

	Describe("sso providers cannot run correctly without proper configs", func() {
		var _sso sso.StrategySelector

		BeforeEach(func() {
			_sso = sso.InitSSO(sso.NewStrategyProxy(
				strategies.NewGithubSSO(githubProcessor),
			))
		})

		It("should prepare login url as intended for github provider", func() {
			githubProcessor.EXPECT().BuildLoginURL().Return(githubInvalidLoginURL, nil).Times(1)
			var url, err = _sso.Algo().Login()

			Expect(err).To(BeNil())
			Expect(url).To(Equal(githubInvalidLoginURL))
		})

		It("should prepare login url as intended for google provider", func() {
			googleProcessor.EXPECT().BuildLoginURL().Return(googleInvalidLoginURL, nil).Times(1)

			googleProvider := sso.NewStrategyProxy(strategies.NewGoogleSSO(googleProcessor))
			_sso.Set(googleProvider)

			var url, err = _sso.Algo().Login()

			Expect(err).To(BeNil())
			Expect(url).To(Equal(googleInvalidLoginURL))
		})

		It("should return github user specific data within provider independent model", func() {
			var expected = &model.Generic{
				Token:         oauthTok.AccessToken,
				Email:         "testmail@github.com",
				VerifiedEmail: true,
			}

			successByProvider(githubProcessor, expected)
			actual, err := _sso.Algo().Register(ctx, oauth.CallbackResponse{})

			Expect(err).To(BeNil())
			Expect(actual).To(Equal(expected))
		})

		It("should return github user specific data within provider independent model", func() {
			var expected = &model.Generic{
				Token:         oauthTok.AccessToken,
				Email:         "testmail@google.com",
				VerifiedEmail: true,
			}

			googleProvider := sso.NewStrategyProxy(strategies.NewGoogleSSO(googleProcessor))
			_sso.Set(googleProvider)

			successByProvider(googleProcessor, expected)
			actual, err := _sso.Algo().Register(ctx, oauth.CallbackResponse{})

			Expect(err).To(BeNil())
			Expect(actual).To(Equal(expected))
		})
	})
})
