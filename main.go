package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/supergiant/supergiant-cli/cmd"
)

func main() {

	app := cli.NewApp()
	app.Name = "supergiant-cli"
	app.Usage = "Powerful control over your supergiants."
	app.Version = Version
	app.Commands = []cli.Command{
		cmd.Create(),
		cmd.Get(),
		cmd.Describe(),
		cmd.Delete(),
		cmd.Select(),
		cmd.Update(),
	}

	app.Run(os.Args)
}
