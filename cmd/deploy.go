package cmd

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/supergiant/supergiant-cli/apictl"
)

func componentDeploy() cli.Command {
	command := cli.Command{
		Name:    "deploy",
		Aliases: []string{"volumes"},
		Usage:   "Deploys a component live.",
		Action: func(c *cli.Context) {
			sg, err := apictl.NewClient("", "", "")
			if err != nil {
				fmt.Println("ERROR:", err)
				os.Exit(5)
			}

			err = sg.DeployComponent(
				c.Args().First(),
				required(c, "app", "Application Name"),
			)
			if err != nil {
				fmt.Println("ERROR:", err)
				os.Exit(5)
			}
			fmt.Println("Success...")
			os.Exit(0)
		},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "app",
				Value: "",
				Usage: "Application Name.",
			},
		},
	}
	return command
}
