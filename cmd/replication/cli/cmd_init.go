package cli

import (
	"github.com/urfave/cli"
)

var initCommand = cli.Command{
	Name:      "init",
	Usage:     "",
	ArgsUsage: "[--home=]",
	Flags:     []cli.Flag{},
	Action:    initRun,
}

func initRun(ctx *cli.Context) error {
	return nil
}
