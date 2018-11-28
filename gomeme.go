package gomeme

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	"golang.org/x/oauth2"
	"log"
)

type Client struct {
	OAuth2Config *oauth2.Config
}

func NewClient(clientID string, clientSecret string, redirectURL string, scopes []string) Client {
	oauth2.RegisterBrokenAuthHeaderProvider("https://apis.jins.com/meme/v1/oauth/token")
	return Client{
		OAuth2Config: &oauth2.Config{
			RedirectURL:  redirectURL,
			ClientID:     clientID,
			ClientSecret: clientSecret,
			Scopes:       scopes,
			Endpoint: oauth2.Endpoint{
				AuthURL:  "https://accounts.jins.com/jp/ja/oauth/authorize",
				TokenURL: "https://apis.jins.com/meme/v1/oauth/token",
			},
		},
	}

}

func (g Client) Exchange(authCode string) {

	token, err := g.OAuth2Config.Exchange(context.Background(), authCode)
	if err != nil {
		log.Fatal(err)
	}

	spew.Dump(token)
	//client := conf.Client(context.Background(), token)
}

func (g Client) GetAuthCode() string {
	url := g.OAuth2Config.AuthCodeURL("state", oauth2.SetAuthURLParam("service_id", "meme"))
	return url
}
