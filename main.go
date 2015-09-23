package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/sorotp/clirescue/trackerapi"
)

func main() {
	app := cli.NewApp()

	app.Name = "clirescue"
	app.Usage = "CLI tool to talk to the Pivotal Tracker's API"

	app.Commands = []cli.Command{
		{
			Name:  "me",
			Usage: "prints out Tracker's representation of your account",
			Action: func(c *cli.Context) {
				trackerapi.Me()
			},
		},
	}

	app.Run(os.Args)
}
