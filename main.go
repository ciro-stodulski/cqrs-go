package main

import (
	app "cqrs-go/cmd/main"
	"log"
)

func main() {
	log.Default().Print("Starting app")

	app.Start()
}
