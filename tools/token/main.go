package main

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"gomeme"
	"log"
	"net/http"
	"os"
)

var meme gomeme.Gomeme

func main() {
	initConfig("config")
	clientID := viper.GetString("Jins.ClientID")
	clientSecret := viper.GetString("Jins.ClientSecret")
	RedirectURL := viper.GetString("Jins.RedirectURL")
	scopes := viper.GetStringSlice("Jins.Scopes")

	meme = gomeme.NewConfig(clientID, clientSecret, RedirectURL, scopes)

	token, err := tokenFromFile("token.json")
	if err != nil {
		url := meme.GetAuthCodeURL()
		fmt.Println(url)
		http.HandleFunc("/", handler)
		http.ListenAndServe(":8080", nil)
	}
	client := gomeme.NewClient(meme, token)
	fmt.Println(client)
}

func handler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	code := r.Form.Get("code")
	fmt.Println(code)
	token := meme.Exchange(code)
	saveToken("token.json", token)

	fmt.Fprintf(w, "Hello, World")
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func initConfig(filename string) {
	viper.AddConfigPath(".")
	viper.SetConfigName(filename)
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
