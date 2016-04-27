package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/supergiant/supergiant-cli/spacetime"
)

// Hack to fill in an required flags from the user.
func required(c *cli.Context, s string, msg string) string {
	var resp string
	if c.String(s) == "" {
		fmt.Print("Please enter " + msg + ": ")
		fmt.Scanln(&resp)
		if resp == "" {
			fmt.Println("ERROR: Empty response.")
			os.Exit(7)
		}
		return resp
	}
	return c.String(s)
}

func context(contextType string) (string, error) {
	switch contextType {
	case "kube":
		kube, err := spacetime.GetDefaultKube()
		if err != nil {
			return "", err
		}
		return kube.Name, nil

	}
	return "", errors.New("Context Not Set.")
}

func getApp(c *cli.Context) string {
	var appName string
	args := c.Args().Tail()
	for i, arg := range args {
		if arg == "--app" {
			appName = args[i+1]
		}
	}
	if appName == "" {
		appName = required(c, "app", "Application Name")
	}
	return appName
}
func getComp(c *cli.Context) string {
	var comp string
	args := c.Args().Tail()
	for i, arg := range args {
		if arg == "--comp" {
			comp = args[i+1]
		}
	}
	if comp == "" {
		comp = required(c, "comp", "Component Name")
	}
	return comp
}
func getFile(c *cli.Context) string {
	var fileName string
	args := c.Args().Tail()
	for i, arg := range args {
		if arg == "--file" {
			fileName = args[i+1]
		}
	}
	return fileName
}
