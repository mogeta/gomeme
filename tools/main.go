package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/viper"
	"gomeme"
	"net/http"
)

var client gomeme.Client

func main() {
	initConfig("config")
	clientID := viper.GetString("Jins.ClientID")
	clientSecret := viper.GetString("Jins.ClientSecret")
	RedirectURL := viper.GetString("Jins.RedirectURL")
	scopes := viper.GetStringSlice("Jins.Scopes")

	client = gomeme.NewClient(clientID, clientSecret, RedirectURL, scopes)
	spew.Dump(client)
	url := client.GetAuthCode()

	fmt.Println(url)

	http.HandleFunc("/", handler)
	go http.ListenAndServe(":8080", nil)

	var s string
	fmt.Scan(&s)
	fmt.Println(s)

}

func handler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	code := r.Form.Get("code")
	fmt.Println(code)
	spew.Dump(client)
	client.Exchange(code)
	fmt.Fprintf(w, "Hello, World")
}

func getAuthTokenURL() {

}

func initConfig(filename string) {
	viper.AddConfigPath(".")
	viper.SetConfigName(filename)
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
