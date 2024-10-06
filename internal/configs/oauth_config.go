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
	return &Config{
		GoogleOauthConfig: &oauth2.Config{
			ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
			ClientSecret: os.Getenv("GOOGLE_CLIENT_ID"),
			Endpoint:     google.Endpoint,
			RedirectURL:  os.Getenv("GOOGLE_CLIENT_ID"),
			Scopes: []string{
				"https://www.googleapis.com/auth/userinfo.email",
				"https://www.googleapis.com/auth/userinfo.profile",
			},
		},
	}
}
