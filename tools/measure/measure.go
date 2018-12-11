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
	//client.GetMeasutreData(start, end, "")
	//client.GetEvents(start, end)

	var cursor = ""
	var computedData []gomeme.ComputedDatum
	for {
		measure, err := client.GetMeasutreData(start, end, cursor)
		for key := range measure.ComputedData {
			computedData = append(computedData, measure.ComputedData[key]...)
		}
		if err != nil {
			log.Fatal(err)
			break
		}
		if measure.Cursor == "" {
			break
		}
		cursor = measure.Cursor
	}
	spew.Dump(computedData)
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
