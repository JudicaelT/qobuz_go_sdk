package qobuz_test

import (
	"fmt"
	"testing"

	"github.com/JudicaelT/qobuzsdk"
	"github.com/JudicaelT/qobuzsdk/internal/qobuz"

	"github.com/stretchr/testify/suite"
)

type LoginWithOAuthTestSuite struct {
	*ApiTestSuite[qobuz.TokenExchange]
	oAuthResponseCode      string
	redirectUrl            string
	endpoint               string
	expectedOAuthLoginLink string
}

func TestLoginWithOAuthTestSuite(t *testing.T) {
	oAuthResponseCode := "response_code"
	redirectUrl := "localhost:3000"
	expectedTokenExchange := qobuz.TokenExchange{Token: "token"}
	testSuite := NewApiTestSuite(expectedTokenExchange)
	endpoint := fmt.Sprintf(
		"oauth/callback?code=%s&private_key=%s",
		oAuthResponseCode,
		testSuite.PrivateKey,
	)
	suite.Run(t, &LoginWithOAuthTestSuite{
		ApiTestSuite:      testSuite,
		oAuthResponseCode: oAuthResponseCode,
		redirectUrl:       redirectUrl,
		endpoint:          endpoint,
		expectedOAuthLoginLink: fmt.Sprintf(
			"https://www.qobuz.com/signin/oauth?ext_app_id=%d&redirect_url=%s",
			testSuite.AppId,
			redirectUrl,
		),
	})
}

func (test *LoginWithOAuthTestSuite) TestLoginWithOAuth() {
	// GIVEN an API client that exchanges an oAuth code for a user token
	client := qobuz.NewClient(
		test.MockApiClient(test.endpoint, nil),
		test.AppId,
		test.AppSecret,
		test.PrivateKey,
		test.GetCurrentTimestamp,
	)
	test.False(client.IsAuthenticated())

	// WHEN login in with oAuth and running the OnCodeReceived callback
	var oAuthLoginLink qobuzsdk.OAuthLogin = client.LoginWithOAuth(test.redirectUrl)
	err := oAuthLoginLink.OnCodeReceived(test.oAuthResponseCode)

	// THEN the correct oAuth login link should be returned
	test.Equal(test.expectedOAuthLoginLink, oAuthLoginLink.LoginLink)
	// AND the OnCodeReceived callback return no error
	test.NoError(err)
	// AND the client should be authenticated
	test.True(client.IsAuthenticated())
}

func (test *LoginWithOAuthTestSuite) TestLoginWithOAuth_GetOAuthCallbackFails() {
	// GIVEN an API client that fails to exchanges an oAuth code for a user token
	oAuthCallbackErr := &qobuzsdk.ApiError{Code: 500, Message: "GET oauth/callback failed"}
	client := qobuz.NewClient(
		test.MockApiClient(test.endpoint, oAuthCallbackErr),
		test.AppId,
		test.AppSecret,
		test.PrivateKey,
		test.GetCurrentTimestamp,
	)
	test.False(client.IsAuthenticated())

	// WHEN login in with oAuth and running the OnCodeReceived callback
	oAuthLoginLink := client.LoginWithOAuth(test.redirectUrl)
	err := oAuthLoginLink.OnCodeReceived(test.oAuthResponseCode)

	// THEN the oAuth login link should be generated correctly
	test.Equal(test.expectedOAuthLoginLink, oAuthLoginLink.LoginLink)
	// BUT the OnCodeReceived callback should fail and return oAuthCallbackErr
	test.ErrorIs(err, oAuthCallbackErr)
	// AND the authentication should fail
	test.False(client.IsAuthenticated())
}
