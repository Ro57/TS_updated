package main

import (
	"fmt"
	"os"
	"token-strike/cmd/replication/cli"
)

func main() {
	app := cli.NewApp()
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "[rpl] %v\n", err)
		os.Exit(1)
	}
}
