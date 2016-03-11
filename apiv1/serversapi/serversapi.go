package serversapi

import (
	"fmt"
	"os"

	"github.com/Jeffail/gabs"
	"github.com/codegangsta/cli"
	"github.com/ddliu/go-httpclient"
	"github.com/galexrt/1cloud-cli/apiv1/authapi"
)

// API for /servers/*

// Servers the api
func ServersList(c *cli.Context, apiURL string) {
	httpclient.Defaults(httpclient.Map{
		httpclient.OPT_USERAGENT: "1cloud-cli",
		"Accept-Language":        "en-us",
		"X-TOKEN":                authapi.GetAuthToken(),
	})
	res, srvErr := httpclient.Get(apiURL+"servers", map[string]string{})
	if srvErr != nil {
		fmt.Printf("There was an error contacting the server. Error: %s", srvErr)
		os.Exit(1)
	}

	bodyString, _ := res.ToString()

	jsonParsed, err := gabs.ParseJSON([]byte(bodyString))

	fmt.Printf("Status Code: %d\nError: %v\n", res.StatusCode, err)
	fmt.Println("Response:", jsonParsed.StringIndent("", "  "))
}
