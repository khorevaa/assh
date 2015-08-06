package main

import (
	"os"
	"path"

	"github.com/moul/advanced-ssh-config/vendor/github.com/codegangsta/cli"

	"github.com/moul/advanced-ssh-config/pkg/commands"
	"github.com/moul/advanced-ssh-config/pkg/version"
)

func main() {
	app := cli.NewApp()
	app.Name = path.Base(os.Args[0])
	app.Author = "Manfred Touron"
	app.Email = "https://github.com/moul/advanced-ssh-config"
	app.Version = version.Version + " (" + version.GitCommit + ")"
	app.Usage = "advanced ssh config"

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug, D",
			Usage: "Enable debug mode",
		},
	}

	app.Commands = commands.Commands

	app.Run(os.Args)
}