package main

import (
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"gomeme"
	"log"
	"os"
	"time"
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
		log.Fatal(err)
	}
	client := gomeme.NewClient(meme, token)

	t := time.Now().AddDate(0, 0, -4)
	start, end := createOneDay(t)
	//fmt.Println(start.String())
	//fmt.Println(end.String())
	//client.GetMeasutreData(start, end, "")
	//client.GetEvents()
	summary, err := client.GetSummary(start, end)
	if err != nil {
		log.Fatal(err)
	}
	spew.Dump(summary)
}

func createOneDay(t time.Time) (begin, end time.Time) {

	begin = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
	end = time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, time.Local)
	return begin, end
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

func initConfig(filename string) {
	viper.AddConfigPath(".")
	viper.SetConfigName(filename)
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}