package gomeme

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"net/url"
	"time"
)

var endpoint = oauth2.Endpoint{
	AuthURL:  "https://accounts.jins.com/jp/ja/oauth/authorize",
	TokenURL: "https://apis.jins.com/meme/v1/oauth/token",
}

type Gomeme struct {
	OAuth2Config *oauth2.Config
}

type memeClient struct {
	Client *http.Client
	Token  *oauth2.Token
}

func NewConfig(clientID string, clientSecret string, redirectURL string, scopes []string) Gomeme {
	oauth2.RegisterBrokenAuthHeaderProvider(endpoint.TokenURL)
	return Gomeme{
		OAuth2Config: &oauth2.Config{
			RedirectURL:  redirectURL,
			ClientID:     clientID,
			ClientSecret: clientSecret,
			Scopes:       scopes,
			Endpoint:     endpoint,
		},
	}

}

func (g Gomeme) Exchange(authCode string) *oauth2.Token {

	token, err := g.OAuth2Config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatal(err)
	}
	return token
}

func (g Gomeme) GetAuthCodeURL() string {
	url := g.OAuth2Config.AuthCodeURL("state", oauth2.SetAuthURLParam("service_id", "meme"))
	return url
}

func NewClient(meme Gomeme, token *oauth2.Token) memeClient {
	return memeClient{
		Client: meme.OAuth2Config.Client(context.TODO(), token),
		Token:  token,
	}
}

func (c memeClient) GetMeasutreData(from, to time.Time, cursor string) {
	v := url.Values{}
	v.Add("access_token", c.Token.AccessToken)
	v.Add("action", "getactivity")

	path := fmt.Sprintf("%s?%s", "https://apis.jins.com/meme/v1/users/me/office2/computed_data", v.Encode())
	req, _ := http.NewRequest("GET", path, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token.AccessToken))
	req.Header.Set("Accept", "application/json")

}
