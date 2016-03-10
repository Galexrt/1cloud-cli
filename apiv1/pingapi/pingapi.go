package pingapi

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/ddliu/go-httpclient"
	"github.com/galexrt/1cloud-cli/apiv1/authapi"
)

// Ping the api
func Ping(c *cli.Context, apiURL string) {
	httpclient.Defaults(httpclient.Map{
		httpclient.OPT_USERAGENT: "1cloud-cli",
		"Accept-Language":        "en-us",
	})
	res, err := httpclient.Get(apiURL+"ping", map[string]string{})
	bodyString, _ := res.ToString()
	fmt.Printf("Response Body: %s\nStatus Code: %d\nError: %v", bodyString, res.StatusCode, err)
}

// PingAuth the api
func PingAuth(c *cli.Context, apiURL string) {
	httpclient.Defaults(httpclient.Map{
		httpclient.OPT_USERAGENT: "1cloud-cli",
		"Accept-Language":        "en-us",
		"X-TOKEN":                authapi.GetAuthToken(),
	})
	res, err := httpclient.Get(apiURL+"ping_auth", map[string]string{})
	bodyString, _ := res.ToString()
	fmt.Printf("Response Body: %s\nStatus Code: %d\nError: %v", bodyString, res.StatusCode, err)
}
