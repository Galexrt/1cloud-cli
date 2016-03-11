package authapi

import (
	"fmt"

	"github.com/codegangsta/cli"
)

// AuthToken given api authentication token
var AuthToken string

// SetAuthToken sets the auth token
func SetAuthToken(authToken string) {
	AuthToken = authToken
}

// GetAuthToken returns the auth token if set or from
func GetAuthToken() string {
	if AuthToken == "" {
		// TODO get auth token from file if exists
	}
	return AuthToken
}

// Auth authenticate an user at the API
func Auth(c *cli.Context, apiURL string) {
	// Ask for username and login to acquire the tokens
	fmt.Printf("Your API Token: %s", "--NOPE--")
}
