package cli

import (
	"log"

	"token-strike/internal/database"
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
	var df database.DBRepository
	server, err := replicatorrpc.New(cfg, df)
	if err != nil {
		log.Fatal(err)
	}

	go server.RunGRPCServer()
	server.RunRestServer()
	return nil
}
