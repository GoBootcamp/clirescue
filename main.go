package main

import (
	"os"

	"github.com/GoBootcamp/clirescue/trackerapi"
	"github.com/codegangsta/cli"
)

var (
	app *cli.App
)

func init() {

	app = cli.NewApp()

	app.Name = "clirescue"
	app.Usage = "CLI tool to talk to the Pivotal Tracker's API"

	app.Commands = []cli.Command{
		{
			Name:  "login",
			Usage: "Logs into your pivotal account and caches your credentials.",
			Action: func(c *cli.Context) {
				trackerapi.CacheCredentials()
			},
		},
	}
}

func main() {
	app.Run(os.Args)
}
