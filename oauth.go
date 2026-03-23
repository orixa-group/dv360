package dv360

import (
	"context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/displayvideo/v4"
)

var apiKey = ""
var apiSecret = ""
var token = &oauth2.Token{}

func WithApiKeys(key, secret string) {
	apiKey = key
	apiSecret = secret
}

func WithRefreshToken(refreshToken string) {
	token = &oauth2.Token{
		RefreshToken: refreshToken,
	}
}

func getOauthConfig(redirectUrl string) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     apiKey,
		ClientSecret: apiSecret,
		RedirectURL:  redirectUrl,
		Scopes: []string{
			// Account management
			displayvideo.DisplayVideoScope,
			// Reporting
			displayvideo.DoubleclickbidmanagerScope,
		},
		Endpoint: google.Endpoint,
	}
}

func GetOauthInstallUrl(redirectUrl string) string {
	return getOauthConfig(redirectUrl).
		AuthCodeURL("", oauth2.AccessTypeOffline, oauth2.ApprovalForce)
}

func GetOauthTokens(redirectUrl, authCode string) (*oauth2.Token, error) {
	return getOauthConfig(redirectUrl).
		Exchange(context.Background(), authCode, oauth2.AccessTypeOffline)
}
