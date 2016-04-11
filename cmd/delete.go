package cmd

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/supergiant/supergiant-cli/apictl"
	"github.com/supergiant/supergiant-cli/core"
	"github.com/supergiant/supergiant-cli/spacetime"
)

// Delete contains subcommands to send information about a resource to
//
func Delete() cli.Command {
	command := cli.Command{
		Name:    "delete",
		Aliases: []string{"d"},
		Usage:   "Supergiant core application control.",
		Subcommands: []cli.Command{

			// Delete Subcommands
			//Destroy a app
			{
				Name:    "application",
				Aliases: []string{"apps", "app", "applications"},
				Usage:   "Destroys Supergaint applications. Aliases: \"apps\", \"app\", \"applications\"",
				Action: func(c *cli.Context) {
					err := apictl.DestroyApp(c.Args().First())
					if err != nil {
						fmt.Println("ERROR:", err)
						os.Exit(5)
					}
					fmt.Println("Success...")
				},
			},
			//Destroy a Component
			{
				Name:    "component",
				Aliases: []string{"components", "comp"},
				Usage:   "Destroys Supergiant application components. Aliases: \"components\", \"comp\"",
				Subcommands: []cli.Command{
					volumeDelete(),
					mountDelete(),
					portDelete(),
					envDelete(),
					commandDelete(),
					containerDelete(),
				},
				Action: func(c *cli.Context) {
					err := apictl.DestroyComponent(c.Args().First(), required(c, "app", "Supergiant App Name"))
					if err != nil {
						fmt.Println("ERROR:", err)
						os.Exit(5)
					}
					fmt.Println("Success...")
				},
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "app",
						Value: "",
						Usage: "Name to assign to the new kubernmetes cluster.",
					},
				},
			},
			//Destroy an Entrypoint
			{
				Name:    "entrypoint",
				Aliases: []string{"entrypoints", "entry", "loadbalancer", "lb"},
				Usage:   "Destroys Supergiant entrypoints. Aliases: \"entrypoints\", \"entry\", \"loadbalancer\", \"lb\"",
				Action: func(c *cli.Context) {
					err := apictl.DestroyEntryPoint(c.Args().First())
					if err != nil {
						fmt.Println("ERROR:", err)
						os.Exit(5)
					}
					fmt.Println("Success...")
				},
			},
			//Destroy an spacetime
			{
				Name:    "spacetime",
				Aliases: []string{"spacetimes", "kube", "kubes", "kubernetes"},
				Usage:   "Destroys Supergaint resources. Aliases: \"kube\", \"kubernetes\"",
				Subcommands: []cli.Command{
					providerDestroy(),
				},
				Action: func(c *cli.Context) {
					err := spacetime.DestroyKube(c.Args().First(), c.Bool("force"))
					if err != nil {
						fmt.Println("ERROR:", err)
						os.Exit(7)
					}
					fmt.Println("Success...")
				},
			},
			//Destroy a core.
			{
				Name:    "core",
				Aliases: []string{"cores", "sg", "supergiant"},
				Usage:   "Destroys Supergiant core.",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "kube",
						Value: "",
						Usage: "Name to assign to the new kubernmetes cluster.",
					},
				},
				Action: func(c *cli.Context) {
					var kube string
					// check for context.
					if c.String("kube") == "" {
						// if flag not set check.
						context, err := context("kube")
						if err != nil {
							fmt.Println("ERROR:", err)
							os.Exit(5)
						}
						kube = context
						// if context not set, require.
						if kube == "" {
							kube = required(c, "kube", "Kubenretes Cluster Context")
						}
					}

					err := sgcore.DestroySGCore(kube)
					if err != nil {
						fmt.Println("Core Install failed: ", err)
						os.Exit(5)
					}
					fmt.Println("Core delete Success...")
				},
			},
			// End spacetime actions.
		},
	}

	return command
}
