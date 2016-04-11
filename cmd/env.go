package cmd

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/supergiant/supergiant-cli/apictl"
)

func envCreate() cli.Command {
	command := cli.Command{
		Name:    "env",
		Aliases: []string{"Envs"},
		Usage:   "Create a new component container Env.",
		Action: func(c *cli.Context) {
			release, err := apictl.GetRelease(
				required(c, "app", "Application Name"),
				required(c, "comp", "Component Name"),
			)
			if err != nil {
				fmt.Println("ERROR:", err)
				os.Exit(5)
			}

			err = apictl.CreateEnv(
				release,
				required(c, "name", "Container Name"),
				required(c, "var", "Env Variable"),
				required(c, "value", "Variable Value"),
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
			cli.StringFlag{
				Name:  "var",
				Value: "gp2",
				Usage: "Env variable.",
			},
			cli.StringFlag{
				Name:  "value",
				Value: "",
				Usage: "variable value.",
			},
		},
	}
	return command
}

func envUpdate() cli.Command {
	command := cli.Command{
		Name:    "env",
		Aliases: []string{"Envs"},
		Usage:   "Create a new component container Env.",
		Action: func(c *cli.Context) {
			release, err := apictl.GetRelease(
				required(c, "app", "Application Name"),
				required(c, "comp", "Component Name"),
			)
			if err != nil {
				fmt.Println("ERROR:", err)
				os.Exit(5)
			}

			err = apictl.UpdateEnv(
				release,
				required(c, "name", "Container Name"),
				required(c, "path", "Env Path"),
				required(c, "volume", "Volume Name"),
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
			cli.StringFlag{
				Name:  "var",
				Value: "gp2",
				Usage: "Env variable.",
			},
			cli.StringFlag{
				Name:  "value",
				Value: "",
				Usage: "variable value.",
			},
		},
	}
	return command
}

func envDelete() cli.Command {
	command := cli.Command{
		Name:    "env",
		Aliases: []string{"Envs"},
		Usage:   "Create a new component container Env.",
		Action: func(c *cli.Context) {
			release, err := apictl.GetRelease(
				required(c, "app", "Application Name"),
				required(c, "comp", "Component Name"),
			)
			if err != nil {
				fmt.Println("ERROR:", err)
				os.Exit(5)
			}

			err = apictl.DeleteEnv(
				release,
				required(c, "name", "Container Name"),
				required(c, "var", "Env Variable"),
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
			cli.StringFlag{
				Name:  "var",
				Value: "gp2",
				Usage: "Env Variable",
			},
		},
	}
	return command
}
