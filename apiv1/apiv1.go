package apiv1

import (
	"github.com/codegangsta/cli"
	"github.com/galexrt/1cloud-cli/apiv1/authapi"
	"github.com/galexrt/1cloud-cli/apiv1/pingapi"
)

// APIURL url to api
var APIURL = "https://cloudpanel-api.1and1.com/v1/"

var authToken string

var commands = []cli.Command{
	{
		Name:    "servers",
		Aliases: []string{"server", "srv"},
		Usage:   "add or list servers ",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "token",
				Value:       "",
				Usage:       "token",
				Destination: &authToken,
			},
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
	{
		Name: "images",
		Action: func(c *cli.Context) {
			println("images")
		},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "image_id",
				Value: "",
				Usage: "image_id",
			},
		},
	},
	{
		Name: "shared_storages",
		Action: func(c *cli.Context) {
			println("shared_storages")
		},
		Subcommands: []cli.Command{
			{
				Name: "servers",
				Action: func(c *cli.Context) {
					println("servers")
				},
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "server_id",
						Value: "",
						Usage: "server_id",
					},
				},
			},
			{
				Name: "access",
				Action: func(c *cli.Context) {
					println("access")
				},
			},
		},
	},
	{
		Name: "firewall_policies",
		Action: func(c *cli.Context) {
			println("firewall_policies")
		},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "firewall_id",
				Value: "",
				Usage: "firewall_id",
			},
		},
		Subcommands: []cli.Command{
			{
				Name: "server_ips",
				Action: func(c *cli.Context) {
					println("server_ips")
				},
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "server_ip",
						Value: "",
						Usage: "server_ip",
					},
				},
			},
			{
				Name: "rules",
				Action: func(c *cli.Context) {
					println("rules")
				},
			},
		},
	},
	{
		Name: "load_balancers",
		Action: func(c *cli.Context) {
			println("load_balancers")
		},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "load_balancer_id",
				Value: "",
				Usage: "load_balancer_id",
			},
		},
		Subcommands: []cli.Command{
			{
				Name: "server_ips",
				Action: func(c *cli.Context) {
					println("server_ips")
				},
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "server_ip",
						Value: "",
						Usage: "server_ip",
					},
				},
			},
			{
				Name: "rules",
				Action: func(c *cli.Context) {
					println("rules")
				},
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "rule_id",
						Value: "",
						Usage: "rule_id",
					},
				},
			},
		},
	},
	{
		Name: "public_ips",
		Action: func(c *cli.Context) {
			println("public_ips")
		},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "ip_id",
				Value: "",
				Usage: "ip_id",
			},
		},
	},
	{
		Name: "private_networks",
		Action: func(c *cli.Context) {
			println("private_networks")
		},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "private_network_id",
				Value: "",
				Usage: "private_network_id",
			},
		},
		Subcommands: []cli.Command{
			{
				Name: "servers",
				Action: func(c *cli.Context) {
					println("servers")
				},
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "server_id",
						Value: "",
						Usage: "server_id",
					},
				},
			},
		},
	},
	{
		Name:    "vpns",
		Aliases: []string{"vpn"},
		Action: func(c *cli.Context) {
			println("vpns")
		},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "vpn_id",
				Value: "",
				Usage: "vpn_id",
			},
		},
		Subcommands: []cli.Command{
			{
				Name: "configuration_file",
				Action: func(c *cli.Context) {
					println("configuration_file")
				},
			},
		},
	},
	{
		Name: "monitoring_center",
		Action: func(c *cli.Context) {
			println("monitoring_center")
		},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "server_id",
				Value: "",
				Usage: "server_id",
			},
		},
	},
	{
		Name: "monitoring_policies",
		Action: func(c *cli.Context) {
			println("monitoring_policies")
		},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "monitoring_policy_id",
				Value: "",
				Usage: "monitoring_policy_id",
			},
		},
		Subcommands: []cli.Command{
			{
				Name: "ports",
				Action: func(c *cli.Context) {
					println("ports")
				},
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "port_id",
						Value: "",
						Usage: "port_id",
					},
				},
			},
			{
				Name: "processes",
				Action: func(c *cli.Context) {
					println("processes")
				},
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "process_id",
						Value: "",
						Usage: "process_id",
					},
				},
			},
			{
				Name: "servers",
				Action: func(c *cli.Context) {
					println("servers")
				},
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "server_id",
						Value: "",
						Usage: "server_id",
					},
				},
			},
		},
	},
	{
		Name: "logs",
		Action: func(c *cli.Context) {
			println("logs")
		},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "log_id",
				Value: "",
				Usage: "log_id",
			},
		},
	},
	{
		Name: "users",
		Action: func(c *cli.Context) {
			println("users")
		},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "user_id",
				Value: "",
				Usage: "user_id",
			},
		},
		Subcommands: []cli.Command{
			{
				Name: "api",
				Action: func(c *cli.Context) {
					println("api")
				},
				Subcommands: []cli.Command{
					{
						Name: "key",
						Action: func(c *cli.Context) {
							println("key")
						},
					},
					{
						Name: "ips",
						Action: func(c *cli.Context) {
							println("ips")
						},
						Flags: []cli.Flag{
							cli.StringFlag{
								Name:  "ip",
								Value: "",
								Usage: "ip",
							},
						},
					},
				},
			},
			{
				Name: "current_user_permissions",
				Action: func(c *cli.Context) {
					println("current_user_permissions")
				},
			},
		},
	},
	{
		Name: "roles",
		Action: func(c *cli.Context) {
			println("roles")
		},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "role_id",
				Value: "",
				Usage: "role_id",
			},
		},
		Subcommands: []cli.Command{
			{
				Name: "permissions",
				Action: func(c *cli.Context) {
					println("permissions")
				},
			},
			{
				Name: "users",
				Action: func(c *cli.Context) {
					println("users")
				},
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "user_id",
						Value: "",
						Usage: "user_id",
					},
				},
			},
			{
				Name: "clone",
				Action: func(c *cli.Context) {
					println("clone")
				},
			},
		},
	},
	{
		Name: "usages",
		Action: func(c *cli.Context) {
			println("usages")
		},
	},
	{
		Name: "server_appliances",
		Action: func(c *cli.Context) {
			println("server_appliances")
		},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "id",
				Value: "",
				Usage: "id",
			},
		},
	},
	{
		Name: "dvd_isos",
		Action: func(c *cli.Context) {
			println("dvd_isos")
		},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "id",
				Value: "",
				Usage: "id",
			},
		},
	},
	{
		Name: "ping",
		Action: func(c *cli.Context) {
			pingapi.Ping(c, APIURL)
		},
	},
	{
		Name: "ping_auth",
		Action: func(c *cli.Context) {
			pingapi.PingAuth(c, APIURL)
		},
	},
	{
		Name: "datacenterse",
		Action: func(c *cli.Context) {
			println("datacenterse")
		},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "datacenter_id",
				Value: "",
				Usage: "datacenter_id",
			},
		},
	},
	{
		Name: "auth",
		Action: func(c *cli.Context) {
			authapi.SetAuthToken(authToken)
			authapi.Auth(c, APIURL)
		},
	},
}

// GetAPICommands returns the command "routes"
func GetAPICommands() []cli.Command {
	return commands
}

// TranslateRespCode translatees the response code
func TranslateRespCode(code int) string {
	switch code {
	case 200:

	}
	return ""
}
