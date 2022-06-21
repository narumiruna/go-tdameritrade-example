package client

import (
	"context"
	"fmt"
	"os"

	"github.com/zricethezav/go-tdameritrade"
	"golang.org/x/oauth2"
)

const tokenURL = "https://api.tdameritrade.com/v1/oauth2/token"
const redirectURL = "https://127.0.0.1"

func New(ctx context.Context) (*tdameritrade.Client, error) {
	consumerKey := os.Getenv("TD_CONSUMER_KEY")
	if consumerKey == "" {
		return nil, fmt.Errorf("TD_CONSUMER_KEY is not set")
	}

	refreshToken := os.Getenv("TD_REFRESH_TOKEN")
	if refreshToken == "" {
		return nil, fmt.Errorf("TD_REFRESH_TOKEN is not set")
	}

	config := oauth2.Config{
		ClientID: consumerKey,
		Endpoint: oauth2.Endpoint{
			TokenURL: tokenURL,
		},
		RedirectURL: redirectURL,
	}

	token := &oauth2.Token{
		RefreshToken: refreshToken,
	}

	httpClient := config.Client(ctx, token)

	return tdameritrade.NewClient(httpClient)
}
