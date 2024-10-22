package configs

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Config struct {
	GoogleOauthConfig *oauth2.Config
}

func NewGoogleConfig() *Config {
	cfg := &Config{
		GoogleOauthConfig: &oauth2.Config{
			Endpoint: google.Endpoint,
			Scopes: []string{
				"https://www.googleapis.com/auth/userinfo.email",
				"https://www.googleapis.com/auth/userinfo.profile",
			},
		},
	}
	cfg.GoogleOauthConfig.ClientID = os.Getenv("GOOGLE_CLIENT_ID")
	cfg.GoogleOauthConfig.ClientSecret = os.Getenv("GOOGLE_CLIENT_SECRET")
	cfg.GoogleOauthConfig.RedirectURL = os.Getenv("GOOGLE_REDIRECT_URL")

	return cfg
}
