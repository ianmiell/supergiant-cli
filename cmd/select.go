package cmd

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/supergiant/supergiant-cli/spacetime"
)

// Select contains subcommands to send information about a resource to
//
func Select() cli.Command {
	command := cli.Command{
		Name:    "select",
		Aliases: []string{"c"},
		Usage:   "Set cli context, avoids having to set context on every command. (You still can though.)",
		Subcommands: []cli.Command{

			// Select Subcommands
			{
				Name:    "application",
				Aliases: []string{"apps", "app", "applications"},
				Usage:   "Lists Supergaint applications. Aliases: \"apps\", \"app\", \"applications\"",
				Action: func(c *cli.Context) {
					fmt.Println("Under Construction")
				},
			},
			//List Components
			{
				Name:    "component",
				Aliases: []string{"components", "comp"},
				Usage:   "Lists Supergiant application components. Aliases: \"components\", \"comp\"",
				Action: func(c *cli.Context) {
					fmt.Println("Under Construction")
				},
			},
			//List Entrypoints
			{
				Name:    "entrypoint",
				Aliases: []string{"entrypoints", "entry", "loadbalancer", "lb"},
				Usage:   "Lists Supergiant entrypoints. Aliases: \"entrypoints\", \"entry\", \"loadbalancer\", \"lb\"",
				Action: func(c *cli.Context) {
					fmt.Println("Under Construction")
				},
			},
			//List Spacetime
			{
				Name:    "spacetime",
				Aliases: []string{"spacetimes", "kube", "kubes", "kubernetes"},
				Usage:   "Lists Supergaint resources. Aliases: \"kube\", \"kubernetes\"",
				Subcommands: []cli.Command{
					providerContext(),
				},
				Action: func(c *cli.Context) {
					err := spacetime.SetDefaultKube(c.Args().First())
					if err != nil {
						fmt.Println("ERROR:", err)
					}
					fmt.Println("Success...")
				},
			},
			// End spacetime actions.
		},
	}

	return command
}
