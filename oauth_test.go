package dv360

import (
	"os"
	"testing"
)

func initTestSession() {
	WithApiKeys(os.Getenv("GOOGLE_CLIENT_ID"), os.Getenv("GOOGLE_CLIENT_SECRET"))
	WithRefreshToken(os.Getenv("GOOGLE_REFRESH_TOKEN"))
}

func TestGetOauthTokens(t *testing.T) {
	initTestSession()

	redirectUrl := os.Getenv("GOOGLE_OAUTH_REDIRECT_URL")

	t.Logf("OAuth install url:\n%s\n\n", GetOauthInstallUrl(redirectUrl))

	if authCode := os.Getenv("GOOGLE_OAUTH_AUTH_CODE"); authCode != "" {
		if token, err := GetOauthTokens(redirectUrl, os.Getenv("GOOGLE_OAUTH_CODE")); nil != err {
			t.Error(err)
		} else if refresh := token.RefreshToken; len(refresh) == 0 {
			t.Error("refresh token is empty")
		} else {
			t.Logf("Refresh token: %s", refresh)
		}
	} else {
		t.Log("Token exchange skipped (missing GOOGLE_OAUTH_CODE en variable)")
	}
}
