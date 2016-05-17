package cmd

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/supergiant/supergiant-cli/apictl"
)

// Describe contains subcommands to send information about a resource to
// stdout.
func Describe() cli.Command {
	command := cli.Command{
		Name:    "describe",
		Aliases: []string{"d"},
		Usage:   "Describe a Supergiant resource.",
		Subcommands: []cli.Command{

			// Describe Subcommands
			//List Apps
			{
				Name:    "application",
				Aliases: []string{"apps", "app", "applications"},
				Usage:   "Describes Supergaint applications. Aliases: \"apps\", \"app\", \"applications\"",
				Action: func(c *cli.Context) {
					sg, err := apictl.NewClient("", "", "")
					if err != nil {
						fmt.Println("ERROR:", err)
						os.Exit(5)
					}

					err = sg.ListApps("")
					if err != nil {
						fmt.Println("ERROR:", err)
						os.Exit(5)
					}
				},
			},
			//List Components
			{
				Name:    "component",
				Aliases: []string{"components", "comp"},
				Usage:   "Describes Supergiant application components. Aliases: \"components\", \"comp\"",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "app",
						Value: "",
						Usage: "Application context.",
					},
				},
				Action: func(c *cli.Context) {
					sg, err := apictl.NewClient("", "", "")
					if err != nil {
						fmt.Println("ERROR:", err)
						os.Exit(5)
					}

					err = sg.ComponentDetails(
						getApp(c),
						c.Args().First(),
					)
					if err != nil {
						fmt.Println("ERROR:", err)
						os.Exit(5)
					}
				},
			},
			//List Entrypoints
			{
				Name:    "entrypoint",
				Aliases: []string{"entrypoints", "entry", "loadbalancer", "lb"},
				Usage:   "Describes Supergiant entrypoints. Aliases: \"entrypoints\", \"entry\", \"loadbalancer\", \"lb\"",
				Action: func(c *cli.Context) {
					sg, err := apictl.NewClient("", "", "")
					if err != nil {
						fmt.Println("ERROR:", err)
						os.Exit(5)
					}

					err = sg.ListEntryPoints(c.Args().First())
					if err != nil {
						fmt.Println("ERROR:", err)
						os.Exit(5)
					}
				},
			},
			// End spacetime actions.
		},
	}

	return command
}
