package cmd

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/supergiant/supergiant-cli/spacetime"
)

func spacetimeInport() cli.Command {

	command := cli.Command{
		Name:  "import",
		Usage: "Import an already existing kubernetes cluster into spacetime.",
		Action: func(c *cli.Context) {
			var provider string
			if c.String("provider") == "" {
				spacetime.ListProvider("")
				provider = required(c, provider, "Provider Name")
			}
			err := spacetime.ImportKube(
				required(c, "name", "Kubernetes Cluster Name"),
				required(c, "ip", "Kubernetes Cluster IP Address"),
				required(c, "user", "Kubernetes Cluster User-Name"),
				required(c, "pass", "Kubernetes Cluster Password"),
				required(c, "region", "AWS Region Name"),
				required(c, "az", "AWS Availability Zone"),
				provider,
			)
			if err != nil {
				fmt.Println("ERROR:", err)
			}
			fmt.Println("Success...")
		},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "ip",
				Value: "",
				Usage: "IP Address of the kubernetes cluster you would like to import.",
			},
			cli.StringFlag{
				Name:  "user",
				Value: "",
				Usage: "User name of the kubernetes cluster you would like to import.",
			},
			cli.StringFlag{
				Name:  "pass",
				Value: "",
				Usage: "Password of the kubernetes cluster you would like to import.",
			},
			cli.StringFlag{
				Name:  "name",
				Value: "",
				Usage: "Cluster name of the kubernetes cluster you would like to import.",
			},
			cli.StringFlag{
				Name:  "region",
				Value: "",
				Usage: "AWS Region of the kubernetes cluster you would like to import.",
			},
			cli.StringFlag{
				Name:  "az",
				Value: "",
				Usage: "AWS Avilability Zone of the kubernetes cluster you would like to import.",
			},
		},
	}

	return command
}
