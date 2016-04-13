package cmd

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/supergiant/supergiant-cli/spacetime"
)

func providerCreate() cli.Command {
	command := cli.Command{
		Name:    "provider",
		Aliases: []string{"providers"},
		Usage:   "Create a new spacetime provider.",
		Action: func(c *cli.Context) {
			if c.Bool("rebuild") {
				err := spacetime.RebuildProvider(
					required(c, "name", "Spacetime Provider Name"),
					c.Bool("verbose"),
				)

				if err != nil {
					fmt.Println("ERROR:", err)
					os.Exit(5)
				}

				if !c.Bool("verbose") {
					fmt.Println("Rebuild submitted...")
					os.Exit(0)
				}

				fmt.Println("Rebuild successful.")
				os.Exit(0)
			}

			err := spacetime.NewProvider(
				required(c, "name", "Spacetime Provider Name"),
				required(c, "access-key", "AWS Access Key"),
				required(c, "secret-key", "AWS Secret Key"),
				required(c, "provider-service", "Host Service"),
				true,
			)
			if err != nil {
				fmt.Println("ERROR:", err)
				os.Exit(5)
			}
		},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "name",
				Value: "",
				Usage: "Name to be used for provider.",
			},
			cli.StringFlag{
				Name:  "access-key",
				Value: "",
				Usage: "AWS Access Key.",
			},
			cli.StringFlag{
				Name:  "secret-key",
				Value: "",
				Usage: "AWS Secret Key.",
			},
			cli.StringFlag{
				Name:  "provider-service, ps",
				Value: "aws",
				Usage: "The cloud service to use. (Only AWS currently supported.)",
			},
			cli.BoolFlag{
				Name:  "rebuild",
				Usage: "Rebuild an exsisting provider. (Requires --name)",
			},
			cli.BoolFlag{
				Name:  "verbose",
				Usage: "Verbose output.",
			},
		},
	}
	return command
}

func providerGet() cli.Command {
	command := cli.Command{
		Name:    "provider",
		Aliases: []string{"providers"},
		Usage:   "Lists spacetime providers.",
		Action: func(c *cli.Context) {
			err := spacetime.ListProvider(c.Args().First())
			if err != nil {
				fmt.Println("ERROR:", err)
				os.Exit(5)
			}
		},
	}
	return command
}

func providerDestroy() cli.Command {

	command := cli.Command{
		Name:    "provider",
		Aliases: []string{"providers"},
		Usage:   "Destroy a spacetime provider. Will fail if provider in use.",
		Action: func(c *cli.Context) {
			err := spacetime.DeleteProvider(c.Args().First())
			if err != nil {
				fmt.Println("ERROR:", err)
				os.Exit(5)
			}
			fmt.Println("Success...")
		},
	}

	return command
}

func providerContext() cli.Command {
	command := cli.Command{
		Name:    "provider",
		Aliases: []string{"providers"},
		Usage:   "Sets a provider context for supergiant commands.",
		Action: func(c *cli.Context) {

		},
	}
	return command
}
