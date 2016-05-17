package cmd

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/supergiant/supergiant-cli/apictl"
)

func mountCreate() cli.Command {
	command := cli.Command{
		Name:    "mount",
		Aliases: []string{"mounts"},
		Usage:   "Create a new component container mount.",
		Action: func(c *cli.Context) {
			sg, err := apictl.NewClient("", "", "")
			if err != nil {
				fmt.Println("ERROR:", err)
				os.Exit(5)
			}

			release, err := sg.GetRelease(
				required(c, "app", "Application Name"),
				required(c, "comp", "Component Name"),
			)
			if err != nil {
				fmt.Println("ERROR:", err)
				os.Exit(5)
			}

			err = apictl.CreateMount(
				release,
				required(c, "name", "Container Name"),
				required(c, "path", "Mount Path"),
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
				Name:  "path",
				Value: "gp2",
				Usage: "Mount Point.",
			},
			cli.StringFlag{
				Name:  "volume",
				Value: "",
				Usage: "Volume to use with this mount.",
			},
		},
	}
	return command
}

func mountUpdate() cli.Command {
	command := cli.Command{
		Name:    "mount",
		Aliases: []string{"mounts"},
		Usage:   "Create a new component container mount.",
		Action: func(c *cli.Context) {
			sg, err := apictl.NewClient("", "", "")
			if err != nil {
				fmt.Println("ERROR:", err)
				os.Exit(5)
			}

			release, err := sg.GetRelease(
				required(c, "app", "Application Name"),
				required(c, "comp", "Component Name"),
			)
			if err != nil {
				fmt.Println("ERROR:", err)
				os.Exit(5)
			}

			err = apictl.UpdateMount(
				release,
				required(c, "name", "Container Name"),
				required(c, "path", "Mount Path"),
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
				Name:  "path",
				Value: "gp2",
				Usage: "Mount Point.",
			},
			cli.StringFlag{
				Name:  "volume",
				Value: "",
				Usage: "Volume to use with this mount.",
			},
		},
	}
	return command
}

func mountDelete() cli.Command {
	command := cli.Command{
		Name:    "mount",
		Aliases: []string{"mounts"},
		Usage:   "Create a new component container mount.",
		Action: func(c *cli.Context) {
			sg, err := apictl.NewClient("", "", "")
			if err != nil {
				fmt.Println("ERROR:", err)
				os.Exit(5)
			}

			release, err := sg.GetRelease(
				required(c, "app", "Application Name"),
				required(c, "comp", "Component Name"),
			)
			if err != nil {
				fmt.Println("ERROR:", err)
				os.Exit(5)
			}

			err = apictl.DeleteMount(
				release,
				required(c, "name", "Container Name"),
				required(c, "path", "Mount Path"),
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
				Name:  "path",
				Value: "gp2",
				Usage: "Mount Point.",
			},
		},
	}
	return command
}
