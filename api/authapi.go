package api

import (
	"github.com/ddliu/go-httpclient"
)

var APIAccessURL = "https://cloudpanel-api.1and1.com/"

var APIV1 = "v1"

func main() {
	httpclient.Defaults(httpclient.Map{
		httpclient.OPT_USERAGENT: "my awsome httpclient",
		"Accept-Language":        "en-us",
	})

	res, err := httpclient.Get("http://google.com/search", map[string]string{
		"q": "news",
	})

	println(res.StatusCode, err)
}

func buildAPIURL() string {
	return APIAccessURL + APIV1
}
