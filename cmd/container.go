package cmd

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/supergiant/supergiant-cli/apictl"
)

func containerCreate() cli.Command {
	command := cli.Command{
		Name:    "container",
		Aliases: []string{"containers"},
		Usage:   "Create a new component container mount.",
		Subcommands: []cli.Command{
			mountCreate(),
		},
		Action: func(c *cli.Context) {
			release, err := apictl.GetRelease(
				required(c, "app", "Application Name"),
				required(c, "comp", "Component Name"),
			)
			if err != nil {
				fmt.Println("ERROR:", err)
				os.Exit(5)
			}

			err = apictl.CreateContainer(
				release,
				required(c, "name", "Container Name"),
				required(c, "image", "Docker image"),
				uint(c.Int("cpu-max")),
				uint(c.Int("cpu-min")),
				uint(c.Int("ram-max")),
				uint(c.Int("ram-min")),
			)
			if err != nil {
				fmt.Println("ERROR:", err)
				os.Exit(5)
			}
			fmt.Println("Success...")
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
				Usage: "Name of container",
			},
			cli.StringFlag{
				Name:  "image",
				Value: "",
				Usage: "Docker image to use with this container.",
			},
			cli.IntFlag{
				Name:  "cpu-max",
				Usage: "CPU Max allocation for this container.",
			},
			cli.IntFlag{
				Name:  "cpu-min",
				Usage: "CPU Min allocation for this container.",
			},
			cli.IntFlag{
				Name:  "ram-max",
				Usage: "RAM Max allocation for this container.",
			},
			cli.IntFlag{
				Name:  "ram-min",
				Usage: "RAM Min allocation for this container.",
			},
		},
	}
	return command
}

func containerUpdate() cli.Command {
	command := cli.Command{
		Name:    "container",
		Aliases: []string{"containers"},
		Usage:   "Create a new component container mount.",
		Subcommands: []cli.Command{
			mountCreate(),
		},
		Action: func(c *cli.Context) {
			release, err := apictl.GetRelease(
				required(c, "app", "Application Name"),
				required(c, "comp", "Component Name"),
			)
			if err != nil {
				fmt.Println("ERROR:", err)
				os.Exit(5)
			}

			err = apictl.UpdateContainer(
				release,
				required(c, "name", "Container Name"),
				required(c, "image", "Docker image"),
				uint(c.Int("cpu-max")),
				uint(c.Int("cpu-min")),
				uint(c.Int("ram-max")),
				uint(c.Int("ram-min")),
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
				Usage: "Name of container",
			},
			cli.StringFlag{
				Name:  "image",
				Value: "",
				Usage: "Docker image to use with this container.",
			},
			cli.IntFlag{
				Name:  "cpu-max",
				Usage: "CPU Max allocation for this container.",
			},
			cli.IntFlag{
				Name:  "cpu-min",
				Usage: "CPU Min allocation for this container.",
			},
			cli.IntFlag{
				Name:  "ram-max",
				Usage: "RAM Max allocation for this container.",
			},
			cli.IntFlag{
				Name:  "ram-min",
				Usage: "RAM Min allocation for this container.",
			},
		},
	}
	return command
}

func containerDelete() cli.Command {
	command := cli.Command{
		Name:    "container",
		Aliases: []string{"containers"},
		Usage:   "Create a new component container mount.",
		Subcommands: []cli.Command{
			mountCreate(),
		},
		Action: func(c *cli.Context) {
			release, err := apictl.GetRelease(
				required(c, "app", "Application Name"),
				required(c, "comp", "Component Name"),
			)
			if err != nil {
				fmt.Println("ERROR:", err)
				os.Exit(5)
			}

			err = apictl.DeleteContainer(
				release,
				required(c, "name", "Container Name"),
			)
			if err != nil {
				fmt.Println("ERROR:", err)
				os.Exit(5)
			}
			fmt.Println("Success...")
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
				Usage: "Name of container",
			},
		},
	}
	return command
}
