package configs

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var GoogleOauthConfig = oauth2.Config{
	ClientID:     "410497696607-bccssf1d7vj0bj35lmp2d5k18l5m0j5p.apps.googleusercontent.com",
	ClientSecret: "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",
	Endpoint:     google.Endpoint,
	RedirectURL:  "http://localhost:3000/auth/google/callback",
	Scopes: []string{
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/userinfo.profile",
	},
}
