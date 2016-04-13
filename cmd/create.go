package cmd

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/supergiant/supergiant-cli/apictl"
	"github.com/supergiant/supergiant-cli/core"
	"github.com/supergiant/supergiant-cli/spacetime"
)

// Create contains subcommands to send information about a resource to
// stdout.
func Create() cli.Command {
	command := cli.Command{
		Name:    "create",
		Aliases: []string{"c"},
		Usage:   "Create a Supergiant resource.",
		Subcommands: []cli.Command{

			// Create Subcommands
			// Create application
			{
				Name:    "application",
				Aliases: []string{"apps", "app", "applications"},
				Usage:   "Creates Supergaint applications. Aliases: \"apps\", \"app\", \"applications\"",
				Action: func(c *cli.Context) {
					err := apictl.CreateApp(c.Args().First())
					if err != nil {
						fmt.Println("ERROR:", err)
						os.Exit(5)
					}
					fmt.Println("Success...")
				},
			},
			// Create component
			{
				Name:    "component",
				Aliases: []string{"components", "comp"},
				Usage:   "Creates Supergiant application components. Aliases: \"components\", \"comp\"",
				Subcommands: []cli.Command{
					volumeCreate(),
					containerCreate(),
					mountCreate(),
					portCreate(),
					envDelete(),
					commandCreate(),
					componentDeploy(),
				},
				Action: func(c *cli.Context) {
					err := apictl.CreateComponent(
						c.Args().First(), // Component Name
						c.String("app"),  // App name
					)
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
			// Create Entrypoint
			{
				Name:    "entrypoint",
				Aliases: []string{"entrypoints", "entry", "loadbalancer", "lb"},
				Usage:   "Creates Supergiant entrypoints. Aliases: \"entrypoints\", \"entry\", \"loadbalancer\", \"lb\"",
				Action: func(c *cli.Context) {
					err := apictl.CreateEntryPoint(
						c.Args().First(), // Entrypoint name
					)
					if err != nil {
						fmt.Println("ERROR:", err)
						os.Exit(5)
					}
					fmt.Println("Success...")
				},
			},
			// Spacetime subcommands
			{
				Name:    "spacetime",
				Aliases: []string{"spacetimes", "kube", "kubes", "kubernetes"},
				Usage:   "Creates Supergaint resources. Aliases: \"kube\", \"kubernetes\"",
				Subcommands: []cli.Command{
					providerCreate(),
					spacetimeInport(),
				},
				Action: func(c *cli.Context) {

					if c.String("retry") != "" {
						err := spacetime.RetryKube(c.String("retry"), true)
						if err != nil {
							fmt.Println("ERROR:", err)
							os.Exit(5)
						}
						os.Exit(0)
					}

					err := spacetime.NewKube(
						required(c, "provider", "Spacetime Provider"),
						required(c, "region", "AWS Region"),
						required(c, "name", "Kube Name"),
						required(c, "user", "Kube UserName"),
						required(c, "pass", "Kube Password"),
						required(c, "avail-zone", "AWS Availability Zone"),
						required(c, "version", "Kubernetes version to use"),
						true,
					)
					if err != nil {
						fmt.Println("ERROR:", err)
						os.Exit(5)
					}

				},
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "name, n",
						Value: "",
						Usage: "Name to assign to the new kubernmetes cluster.",
					},
					cli.StringFlag{
						Name:  "version, kv",
						Value: "1.1.7",
						Usage: "Version of the new kubernmetes cluster.",
					},
					cli.StringFlag{
						Name:   "provider, pvdr",
						Value:  "aws",
						Usage:  "Provider to use while launching kubernetes.",
						EnvVar: "CLOUD_PROVIDER",
					},
					cli.StringFlag{
						Name:   "region, r",
						Value:  "us-east-1",
						Usage:  "Region where kubernetes cluster will live.",
						EnvVar: "CLOUD_REGION",
					},
					cli.StringFlag{
						Name:  "user, u",
						Value: "",
						Usage: "Username to use with your kubermetes cluster.",
					},
					cli.StringFlag{
						Name:  "pass, p",
						Value: "",
						Usage: "Password to use with your kubermetes cluster.",
					},
					cli.StringFlag{
						Name:  "avail-zone, az",
						Value: "us-east-1b",
						Usage: "AWS availability zone where your kubernetes cluster will live.",
					},
					cli.StringFlag{
						Name:  "retry",
						Usage: "Retry a failed install.",
					},
				},
			},

			// Create Entrypoint
			{
				Name:    "core",
				Aliases: []string{"cores", "sg", "supergiant"},
				Usage:   "Creates Supergiant core.",
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
					fmt.Println("Kube context is:", kube)
					err := sgcore.InstallSGCore(kube)
					if err != nil {
						fmt.Println("Core Install failed: ", err)
						os.Exit(5)
					}
					fmt.Println("Core install Success...")
				},
			},

			// End spacetime actions.
		},
	}

	return command
}
