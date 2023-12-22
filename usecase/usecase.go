package usecase

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"net/url"
)

type InitAuthResponse struct {
	Url   string `json:"url"`
	State string `json:"state"`
}

type CallbackResponse struct {
	Id      string `json:"id"`
	Email   string `json:"email"`
	Picture string `json:"picture"`
	Token   struct {
		AccessToken  string `json:"accessToken"`
		RefreshToken string `json:"refreshToken"`
		Expiry       string `json:"Expiry"`
	} `json:"token"`
}

var oauthConfig = &oauth2.Config{
	Scopes:   []string{"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint: google.Endpoint,
}

func InitAuth(ctx context.Context, url, clientID, clientSecret string) (*InitAuthResponse, error) {
	oauthConfig.RedirectURL = url
	oauthConfig.ClientID = clientID
	oauthConfig.ClientSecret = clientSecret

	state := generateState()

	return &InitAuthResponse{
		Url:   oauthConfig.AuthCodeURL(state),
		State: state,
	}, nil
}

func Callback(ctx context.Context, code string) (*CallbackResponse, error) {
	token, err := oauthConfig.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}

	u := "https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + url.QueryEscape(token.AccessToken)
	client := resty.New()

	data, err := client.R().Get(u)
	if err != nil {
		return nil, err
	}

	var res *CallbackResponse

	_ = json.Unmarshal(data.Body(), &res)

	res.Token.AccessToken = token.AccessToken
	res.Token.RefreshToken = token.RefreshToken
	res.Token.Expiry = token.Expiry.String()

	return res, nil
}

func generateState() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)

	return base64.URLEncoding.EncodeToString(b)
}
