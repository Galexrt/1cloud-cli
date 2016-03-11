package authapi

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
