package apiv1

import (
	"github.com/codegangsta/cli"
)

var commands = []cli.Command{
	{
		Name:    "servers",
		Aliases: []string{"net"},
		Usage:   "add or list servers ",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "server_id",
				Value: "",
				Usage: "server_id",
			},
		},
		Subcommands: []cli.Command{
			{
				Name:    "list",
				Aliases: []string{"ls"},
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
			{
				Name: "fixed_instance_sizes",
				Action: func(c *cli.Context) {
					println("fixed_instance_sizes", c.Args().First())
				},
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "fixed_instance_size_id",
						Value: "",
						Usage: "fixed_instance_size_id",
					},
				},
			},
			{
				Name: "hardware",
				Action: func(c *cli.Context) {
					println("hardware", c.Args().First())
				},
				Subcommands: []cli.Command{
					{
						Name: "hdds",
						Action: func(c *cli.Context) {
							println("hdds", c.Args().First())
						},
						Flags: []cli.Flag{
							cli.StringFlag{
								Name:  "hdd_id",
								Value: "",
								Usage: "hdd_id",
							},
						},
					},
				},
			},
			{
				Name: "image",
				Action: func(c *cli.Context) {
					println("image", c.Args().First())
				},
			},
			{
				Name: "ips",
				Action: func(c *cli.Context) {
					println("ips", c.Args().First())
				},
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "ip_id",
						Value: "",
						Usage: "ip_id",
					},
				},
				Subcommands: []cli.Command{
					{
						Name: "firewall_policy",
						Action: func(c *cli.Context) {
							println("firewall_policy", c.Args().First())
						},
					},
					{
						Name: "load_balancers",
						Action: func(c *cli.Context) {
							println("load_balancers", c.Args().First())
						},
						Flags: []cli.Flag{
							cli.StringFlag{
								Name:  "load_balancer_id",
								Value: "",
								Usage: "load_balancer_id",
							},
						},
					},
				},
			},
			{
				Name: "status",
				Action: func(c *cli.Context) {
					println("status", c.Args().First())
				},
			},
			{
				Name: "dvd",
				Action: func(c *cli.Context) {
					println("dvd", c.Args().First())
				},
			},
			{
				Name: "private_networks",
				Action: func(c *cli.Context) {
					println("private_ networks", c.Args().First())
				},
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "private_network_id",
						Value: "",
						Usage: "private_network_id",
					},
				},
			},
			{
				Name: "snapshots",
				Action: func(c *cli.Context) {
					println("snapshots", c.Args().First())
				},
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "snapshot_id",
						Value: "",
						Usage: "snapshot_id",
					},
				},
			},
			{
				Name: "clone",
				Action: func(c *cli.Context) {
					println("clone", c.Args().First())
				},
			},
		},
	},
}

// GetAPICommands returns the command "routes"
func GetAPICommands() []cli.Command {
	return commands
}
