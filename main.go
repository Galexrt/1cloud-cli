package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/galexrt/1cloud-cli/apiv1"
)

type cmdLineOpts struct {
	help    bool
	version bool
}

var opts cmdLineOpts

func showHelp() {
	fmt.Printf("%s [command] [args]", os.Args[0])
	flag.PrintDefaults()
	os.Exit(0)
}

func main() {
	flag.Parse()

	app := cli.NewApp()
	app.Name = "1cloud-cli"
	app.Usage = "access to 1&1 cloud api"
	app.EnableBashCompletion = true

	app.Commands = apiv1.GetAPICommands()

	app.Run(os.Args)
}
