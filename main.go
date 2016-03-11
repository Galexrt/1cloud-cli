package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/galexrt/1cloud-cli/apiv1"
	"github.com/galexrt/1cloud-cli/apiv1/authapi"
)

func main() {
	app := cli.NewApp()
	app.Name = "1cloud-cli"
	app.Usage = "access to 1&1 cloud api"
	app.EnableBashCompletion = true
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "token",
			Value:       "",
			Usage:       "token",
			Destination: &authapi.AuthToken,
		},
	}
	app.Commands = apiv1.GetAPICommands()

	app.Run(os.Args)
}
