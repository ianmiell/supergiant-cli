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
				Aliases: []string{"components", "comp", "comps"},
				Usage:   "Lists Supergiant application components. Aliases: \"components\", \"comp\"",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "app",
						Value: "",
						Usage: "Application Name.",
					},
					cli.StringFlag{
						Name:  "output, o",
						Value: "",
						Usage: "Output component in json/yaml format.",
					},
					cli.BoolFlag{
						Name:  "example",
						Usage: "Output example YAML/JSON file configs.",
					},
				},
				Action: func(c *cli.Context) {
					sg, err := apictl.NewClient("", "", "")
					if err != nil {
						fmt.Println("ERROR:", err)
						os.Exit(5)
					}

					// Print Example info
					if c.Bool("example") {
						err := apictl.GetReleaseExample()
						if err != nil {
							fmt.Println("ERROR:", err)
							os.Exit(5)
						}
						os.Exit(0)
					}

					// Output in custom format.
					if c.String("output") != "" {
						if c.Args().First() == "" {
							fmt.Println("Specify a component name. \"supergiant get comp <foo> -o json\"")
							os.Exit(5)
						}
						err = sg.ListCompenentinFormat(
							c.String("output"),
							required(c, "app", "Applcation Name"),
							c.Args().First(),
						)
						if err != nil {
							fmt.Println("ERROR:", err)
							os.Exit(5)
						}
						os.Exit(0)
					}

					// Show all components.
					if c.String("app") == "" {
						err := sg.ListAllComponents()
						if err != nil {
							fmt.Println("ERROR:", err)
							os.Exit(5)
						}
						os.Exit(0)
					}

					// Show only components for a specified app.
					err = sg.ListComponents(c.String("app"))
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
