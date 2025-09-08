package main

import (
	"log"
	"synapse-ai/internal/server"
)

func main() {
	svr := server.CreateServer(":8020")

	if err := svr.Run(); err != nil {
		log.Fatal(err)
	}
}
