package main

import (
	"log"

	"token-strike/server/replicatorrpc"
)

func main() {
	host := ":5300"
	someDomain := "http://some.com"

	server, err := replicatorrpc.New(host, someDomain)
	if err != nil {
		log.Fatal(err)
	}

	server.RunGRPCServer(host)

}
