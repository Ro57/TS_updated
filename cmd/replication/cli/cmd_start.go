package cli

import (
	"log"

	"token-strike/server/replicatorrpc"

	"github.com/urfave/cli"
)

var startCommand = cli.Command{
	Name:      "start",
	Usage:     "",
	ArgsUsage: "[--home=]",
	Flags:     []cli.Flag{},
	Action:    preRunDecorator(startRun),
}

func startRun(ctx *cli.Context) error {
	server, err := replicatorrpc.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	go server.RunGRPCServer()
	server.RunRestServer()
	return nil
}
