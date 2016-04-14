package cmd

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/supergiant/supergiant-cli/apictl"
	"github.com/supergiant/supergiant-cli/spacetime"
)

// Get contains subcommands to send information about a resource to
//
func Get() cli.Command {
	command := cli.Command{
		Name:    "list",
		Aliases: []string{"get"},
		Usage:   "List a Supergiant resource.",
		Subcommands: []cli.Command{

			// Get Subcommands
			// Describe Subcommands
			//List Apps
			{
				Name:    "application",
				Aliases: []string{"apps", "app", "applications"},
				Usage:   "Lists Supergaint applications. Aliases: \"apps\", \"app\", \"applications\"",
				Action: func(c *cli.Context) {
					err := apictl.ListApps("")
					if err != nil {
						fmt.Println("ERROR:", err)
						os.Exit(5)
					}
				},
			},
			//List Components
			{
				Name:    "component",
				Aliases: []string{"components", "comp", "comps"},
				Usage:   "Lists Supergiant application components. Aliases: \"components\", \"comp\"",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "app",
						Value: "",
						Usage: "Name to assign to the new kubernmetes cluster.",
					},
				},
				Action: func(c *cli.Context) {
					if c.String("app") == "" {
						err := apictl.ListAllComponents()
						if err != nil {
							fmt.Println("ERROR:", err)
							os.Exit(5)
						}
						os.Exit(0)
					}

					err := apictl.ListComponents(c.String("app"))
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
				Usage:   "Lists Supergiant entrypoints. Aliases: \"entrypoints\", \"entry\", \"loadbalancer\", \"lb\"",
				Action: func(c *cli.Context) {
					err := apictl.ListEntryPoints(c.Args().First())
					if err != nil {
						fmt.Println("ERROR:", err)
						os.Exit(5)
					}
				},
			},
			//List Spacetime
			{
				Name:    "spacetime",
				Aliases: []string{"spacetimes", "kube", "kubes", "kubernetes"},
				Usage:   "Lists Supergaint resources. Aliases: \"kube\", \"kubernetes\"",
				Subcommands: []cli.Command{
					providerGet(),
				},
				Action: func(c *cli.Context) {
					spacetime.ListKubes(c.Args().First())
				},
			},
			// End spacetime actions.
		},
	}

	return command
}
