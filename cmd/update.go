package cmd

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/supergiant/supergiant-cli/core"
)

// Update contains subcommands to send information about a resource to
//
func Update() cli.Command {
	command := cli.Command{
		Name:    "update",
		Aliases: []string{"u", "upgrade"},
		Usage:   "Update a Supergiant resource.",
		Subcommands: []cli.Command{
			// Update Subcommands
			// Update Comp
			{
				Name:    "component",
				Aliases: []string{"components", "comp"},
				Usage:   "Creates Supergiant application components. Aliases: \"components\", \"comp\"",
				Subcommands: []cli.Command{
					volumeUpdate(),
					mountUpdate(),
					portUpdate(),
					envUpdate(),
					containerUpdate(),
				},
				Action: func(c *cli.Context) {},
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "app",
						Value: "",
						Usage: "Name to assign to the new kubernmetes cluster.",
					},
				},
			},
			// Update core
			{
				Name:    "core",
				Aliases: []string{"cores", "sg", "supergiant"},
				Usage:   "Updates Supergiant core.",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "kube",
						Value: "",
						Usage: "Name of the Kubernetes/Supergiant cluster to update.",
					},
					cli.StringFlag{
						Name:  "api-version",
						Value: "",
						Usage: "Version you would like to update to. (A default latest stable release is set.)",
					},
					cli.StringFlag{
						Name:  "dash-version",
						Value: "",
						Usage: "Version you would like to update to. (A default latest stable release is set.)",
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

					if c.String("api-version") != "" {
						err := sgcore.UpdateSGCore(
							kube,
							c.String("api-version"),
						)
						if err != nil {
							fmt.Println("ERROR:", err)
							os.Exit(5)
						}
					}

					if c.String("dash-version") != "" {
						err := sgcore.UpdateDash(
							kube,
							c.String("dash-version"),
						)
						if err != nil {
							fmt.Println("ERROR:", err)
							os.Exit(5)
						}
					}
					fmt.Println("Success...")
				},
			},

			// End spacetime actions.
		},
	}

	return command
}
