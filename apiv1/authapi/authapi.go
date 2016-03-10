package authapi

import (
	"fmt"

	"github.com/codegangsta/cli"
)

// AuthToken given api authentication token
var AuthToken string

func SetAuthToken(authToken string) {
	AuthToken = authToken
}

func GetAuthToken() string {
	return AuthToken
}

func Auth(c *cli.Context, apiURL string) {
	// Ask for username and login to acquire the tokens
	fmt.Printf("Your API Token: %s", "--NOPE--")
}
