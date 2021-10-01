package cli

// start - started server
// init - initialize config

import (
	"github.com/urfave/cli"
)

func NewApp() *cli.App {
	app := cli.NewApp()
	app.Name = "rpl"
	app.Usage = ""
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "home",
			Usage: "",
		},
	}

	app.Commands = []cli.Command{
		initCommand,
		startCommand,
	}

	return app
}
