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
	Action:    startRun,
}

func startRun(ctx *cli.Context) error {

	host := ":8081"
	someDomain := "http://some.com"

	server, err := replicatorrpc.New(host, someDomain)
	if err != nil {
		log.Fatal(err)
	}

	server.RunGRPCServer(host)
	return nil
}
