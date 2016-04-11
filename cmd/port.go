package cmd

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/supergiant/supergiant-cli/apictl"
)

func portCreate() cli.Command {
	command := cli.Command{
		Name:    "port",
		Aliases: []string{"Ports"},
		Usage:   "Create a new component container Port.",
		Action: func(c *cli.Context) {
			release, err := apictl.GetRelease(
				required(c, "app", "Application Name"),
				required(c, "comp", "Component Name"),
			)
			if err != nil {
				fmt.Println("ERROR:", err)
				os.Exit(5)
			}

			err = apictl.CreatePort(
				release,
				required(c, "name", "Container Name"),
				required(c, "protocol", "Port Protocol"),
				c.Int("port-number"),
				c.Bool("public"),
				c.String("entrypoint"),
				c.Int("external-port-number"),
			)
			if err != nil {
				fmt.Println("ERROR:", err)
				os.Exit(5)
			}
		},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "comp",
				Value: "",
				Usage: "Component Name.",
			},
			cli.StringFlag{
				Name:  "app",
				Value: "",
				Usage: "Application Name.",
			},
			cli.StringFlag{
				Name:  "name",
				Value: "",
				Usage: "Container Name.",
			},
			cli.StringFlag{
				Name:  "protocol",
				Value: "",
				Usage: "Protocol for port.",
			},
			cli.IntFlag{
				Name:  "port-number",
				Usage: "Port number to use with this port.",
			},
			cli.BoolFlag{
				Name:  "public",
				Usage: "Should this port be public?",
			},
			cli.StringFlag{
				Name:  "entrypoint",
				Value: "",
				Usage: "Which entrypoint/loadbalancer should be used with this port.",
			},
			cli.IntFlag{
				Name:  "external-port-number",
				Usage: "(requires public flag) Would you like a specific external port number? if not set, external port will be random.",
			},
		},
	}
	return command
}

func portUpdate() cli.Command {
	command := cli.Command{
		Name:    "port",
		Aliases: []string{"Ports"},
		Usage:   "Create a new component container Port.",
		Action: func(c *cli.Context) {
			release, err := apictl.GetRelease(
				required(c, "app", "Application Name"),
				required(c, "comp", "Component Name"),
			)
			if err != nil {
				fmt.Println("ERROR:", err)
				os.Exit(5)
			}

			err = apictl.UpdatePort(
				release,
				required(c, "name", "Container Name"),
				required(c, "protocol", "Port Protocol"),
				c.Int("port-number"),
				c.Bool("public"),
				c.String("entrypoint"),
				c.Int("external-port-number"),
			)
			if err != nil {
				fmt.Println("ERROR:", err)
				os.Exit(5)
			}
		},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "comp",
				Value: "",
				Usage: "Component Name.",
			},
			cli.StringFlag{
				Name:  "app",
				Value: "",
				Usage: "Application Name.",
			},
			cli.StringFlag{
				Name:  "name",
				Value: "",
				Usage: "Container Name.",
			},
			cli.StringFlag{
				Name:  "protocol",
				Value: "",
				Usage: "Protocol for port.",
			},
			cli.IntFlag{
				Name:  "port-number",
				Usage: "Port number to use with this port.",
			},
			cli.BoolFlag{
				Name:  "public",
				Usage: "Should this port be public?",
			},
			cli.StringFlag{
				Name:  "entrypoint",
				Value: "",
				Usage: "Which entrypoint/loadbalancer should be used with this port.",
			},
			cli.IntFlag{
				Name:  "external-port-number",
				Usage: "(requires public flag) Would you like a specific external port number? if not set, external port will be random.",
			},
		},
	}
	return command
}

func portDelete() cli.Command {
	command := cli.Command{
		Name:    "port",
		Aliases: []string{"Ports"},
		Usage:   "Create a new component container Port.",
		Action: func(c *cli.Context) {
			release, err := apictl.GetRelease(
				required(c, "app", "Application Name"),
				required(c, "comp", "Component Name"),
			)
			if err != nil {
				fmt.Println("ERROR:", err)
				os.Exit(5)
			}

			err = apictl.DeletePort(
				release,
				required(c, "name", "Container Name"),
				c.Int("port-number"),
			)
			if err != nil {
				fmt.Println("ERROR:", err)
				os.Exit(5)
			}
		},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "comp",
				Value: "",
				Usage: "Component Name.",
			},
			cli.StringFlag{
				Name:  "app",
				Value: "",
				Usage: "Application Name.",
			},
			cli.StringFlag{
				Name:  "name",
				Value: "",
				Usage: "Container Name.",
			},
			cli.IntFlag{
				Name:  "port-number",
				Usage: "Port number to use with this port.",
			},
		},
	}
	return command
}
