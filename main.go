package main

import (
	"os"

	"github.com/sourabp/clirescue/trackerapi"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()

	app.Name = "clirescue"
	app.Usage = "CLI tool to talk to the Pivotal Tracker's API"

	app.Commands = []*cli.Command{
		{
			Name:  "me",
			Usage: "prints out Tracker's representation of your account",
			Action: func(c *cli.Context) error {
				return trackerapi.Me()
			},
		},
	}

	app.Run(os.Args)
}
