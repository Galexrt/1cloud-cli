package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/coreos/pkg/flagutil"
	"github.com/galexrt/1u1cloud-cli/apiauth"
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
	if opts.version {
		fmt.Printf("1u1cloud-cli")
		flag.PrintDefaults()
		os.Exit(0)
	}
	if opts.help {
		showHelp()
	}
	flagutil.SetFlagsFromEnv(flag.CommandLine, "1U1CLOUD")

	switch os.Args[1] {
	case "auth":
		apiauth.Auth()
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
	case "help":
		showHelp()
	default:
		fmt.Print("No command given.")
		os.Exit(2)
	}
}
