package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/codegangsta/cli"
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
	app.Name = "boom"
	app.Usage = "make an explosive entrance"
	app.Action = func(c *cli.Context) {
		println("boom! I say!")
	}
	app.EnableBashCompletion = true

	app.Commands = []cli.Command{
		{
			Name:    "servers",
			Aliases: []string{"net"},
			Usage:   "add or list servers ",
			Subcommands: []cli.Command{
				{
					Name: "list",
					Action: func(c *cli.Context) {
						println("server list", c.Args().First())
					},
				},
				{
					Name: "add",
					Action: func(c *cli.Context) {
						println("server list", c.Args().First())
					},
				},
			},
		},
		{
			Name:    "complete",
			Aliases: []string{"c"},
			Usage:   "complete a task on the list",
			Action: func(c *cli.Context) {
				println("completed task: ", c.Args().First())
			},
		},
		{
			Name:    "template",
			Aliases: []string{"r"},
			Usage:   "options for task templates",
			Subcommands: []cli.Command{
				{
					Name:  "add",
					Usage: "add a new template",
					Action: func(c *cli.Context) {
						println("new task template: ", c.Args().First())
					},
				},
				{
					Name:  "remove",
					Usage: "remove an existing template",
					Action: func(c *cli.Context) {
						println("removed task template: ", c.Args().First())
					},
				},
			},
		},
	}

	app.Run(os.Args)
	/*
		case "auth":

		case "servers":

		case "images":

		case "shared_storages":
			fallthrough
		case "sharedstorages":

		case "firewall":
			fallthrough
		case "firewallpolicies":

		case "load_balancer":
			fallthrough
		case "loadbalancer":

		case "ips":
			fallthrough
		case "public_ips":
			fallthrough
		case "publicips":

		case "privatenetworks":
			fallthrough
		case "private_networks":

		case "vpn":
			fallthrough
		case "monitoring_center":
			fallthrough
		case "monitoringcenter":

		case "monitoring_policies":

		case "monitoringpolicies":

		case "logs":

		case "users":

		case "roles":

		case "serverappliances":
			fallthrough
		case "server_appliances":

		case "dvdiso":
			fallthrough
		case "dvd_iso":

			// ping authentication has been ignored here
		case "ping":

		case "pricing":

		case "datacenters":

		case "":
			fallthrough
	*/
}
