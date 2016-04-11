package cmd

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/supergiant/supergiant-cli/apictl"
)

func commandCreate() cli.Command {
	command := cli.Command{
		Name:    "command",
		Aliases: []string{"Commands"},
		Usage:   "Create a new component container Command.",
		Action: func(c *cli.Context) {
			release, err := apictl.GetRelease(
				required(c, "app", "Application Name"),
				required(c, "comp", "Component Name"),
			)
			if err != nil {
				fmt.Println("ERROR:", err)
				os.Exit(5)
			}

			err = apictl.CreateCommand(
				release,
				required(c, "name", "Container Name"),
				c.StringSlice("cmd"),
			)
			if err != nil {
				fmt.Println("ERROR:", err)
				os.Exit(5)
			}
		},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "comp",
				Value: "",
				Usage: "Component Name.",
			},
			cli.StringFlag{
				Name:  "app",
				Value: "",
				Usage: "Application Name.",
			},
			cli.StringFlag{
				Name:  "name",
				Value: "",
				Usage: "Container Name.",
			},
			cli.StringSliceFlag{
				Name:  "cmd",
				Usage: "Command variable.",
			},
		},
	}
	return command
}

func commandDelete() cli.Command {
	command := cli.Command{
		Name:    "command",
		Aliases: []string{"Commands"},
		Usage:   "Create a new component container Command.",
		Action: func(c *cli.Context) {
			release, err := apictl.GetRelease(
				required(c, "app", "Application Name"),
				required(c, "comp", "Component Name"),
			)
			if err != nil {
				fmt.Println("ERROR:", err)
				os.Exit(5)
			}

			err = apictl.DeleteCommand(
				release,
				required(c, "name", "Container Name"),
				c.StringSlice("cmd"),
			)
			if err != nil {
				fmt.Println("ERROR:", err)
				os.Exit(5)
			}
		},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "comp",
				Value: "",
				Usage: "Component Name.",
			},
			cli.StringFlag{
				Name:  "app",
				Value: "",
				Usage: "Application Name.",
			},
			cli.StringFlag{
				Name:  "name",
				Value: "",
				Usage: "Container Name.",
			},
			cli.StringSliceFlag{
				Name:  "cmd",
				Usage: "Command Variable",
			},
		},
	}
	return command
}
