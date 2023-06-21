package reposiroty

import (
	"github.com/04Akaps/Video_Chat_App/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const (
	ScopeEmail   = "https://www.googleapis.com/auth/userinfo.email"
	ScopeProfile = "https://www.googleapis.com/auth/userinfo.profile"
)

type Auth struct {
	OAuthConf *oauth2.Config
}

func NewOAuth(cfg *config.Config) *Auth {

	oAuthConf := &oauth2.Config{
		ClientID:     cfg.GoogleOAuth.ClientID,
		ClientSecret: cfg.GoogleOAuth.ClientSecret,
		RedirectURL:  cfg.GoogleOAuth.CallbackUrl,
		Scopes:       []string{ScopeEmail, ScopeProfile},
		Endpoint:     google.Endpoint,
	}

	r := Auth{
		OAuthConf: oAuthConf,
	}

	return &r
}
