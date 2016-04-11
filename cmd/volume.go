package cmd

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/supergiant/supergiant-cli/apictl"
)

func volumeCreate() cli.Command {
	command := cli.Command{
		Name:    "volume",
		Aliases: []string{"volumes"},
		Usage:   "Create a new component volume.",
		Action: func(c *cli.Context) {
			release, err := apictl.GetRelease(
				required(c, "app", "Application Name"),
				required(c, "comp", "Component Name"),
			)
			if err != nil {
				fmt.Println("ERROR:", err)
				os.Exit(5)
			}

			err = apictl.CreateVolume(
				release,
				required(c, "name", "Volume Name"),
				c.String("type"),
				c.Int("size"),
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
				Usage: "Volume Name.",
			},
			cli.StringFlag{
				Name:  "type",
				Value: "gp2",
				Usage: "Volume type.",
			},
			cli.IntFlag{
				Name:  "size",
				Value: 20,
				Usage: "Volume size.",
			},
		},
	}
	return command
}

func volumeUpdate() cli.Command {
	command := cli.Command{
		Name:    "volume",
		Aliases: []string{"volumes"},
		Usage:   "Create a new component volume.",
		Action: func(c *cli.Context) {
			release, err := apictl.GetRelease(
				required(c, "app", "Application Name"),
				required(c, "comp", "Component Name"),
			)
			if err != nil {
				fmt.Println("ERROR:", err)
				os.Exit(5)
			}

			err = apictl.UpdateVolume(
				release,
				required(c, "name", "Volume Name"),
				c.String("type"),
				c.Int("size"),
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
				Usage: "Volume Name.",
			},
			cli.StringFlag{
				Name:  "type",
				Value: "gp2",
				Usage: "Volume type.",
			},
			cli.IntFlag{
				Name:  "size",
				Value: 20,
				Usage: "Volume size.",
			},
		},
	}
	return command
}

func volumeDelete() cli.Command {
	command := cli.Command{
		Name:    "volume",
		Aliases: []string{"volumes"},
		Usage:   "Create a new component volume.",
		Action: func(c *cli.Context) {
			release, err := apictl.GetRelease(
				required(c, "app", "Application Name"),
				required(c, "comp", "Component Name"),
			)
			if err != nil {
				fmt.Println("ERROR:", err)
				os.Exit(5)
			}

			err = apictl.DeleteVolume(
				release,
				required(c, "name", "Volume Name"),
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
				Usage: "Volume Name.",
			},
		},
	}
	return command
}
