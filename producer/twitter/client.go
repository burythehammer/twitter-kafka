package twitterclient

import (
	"github.com/dghubble/go-twitter/twitter"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"os"
)

func CreateTwitterClient() *twitter.Client {
	details := getTwitterLoginDetails()

	// oauth2 configures a client that uses app credentials to keep a fresh token
	config := &clientcredentials.Config{
		ClientID:     details.ApiKey,
		ClientSecret: details.ApiSecretKey,
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}
	// http.Client will automatically authorize Requests
	httpClient := config.Client(oauth2.NoContext)

	// Twitter client
	client := twitter.NewClient(httpClient)
	return client
}

type twitterLoginDetails struct {
	ApiKey       string
	ApiSecretKey string
	BearerToken  string
}

func getTwitterLoginDetails() twitterLoginDetails {
	return twitterLoginDetails{
		ApiKey:       getApiKey(),
		ApiSecretKey: getApiSecretKey(),
		BearerToken:  getBearerToken(),
	}
}

func getApiKey() string {
	return os.Getenv("TWITTER_API_KEY")
}

func getApiSecretKey() string {
	return os.Getenv("TWITTER_API_SECRET_KEY")
}

func getBearerToken() string {
	return os.Getenv("TWITTER_BEARER_TOKEN")
}
