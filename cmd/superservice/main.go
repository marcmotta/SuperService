// cmd/superservice/main.go
package main

import (
	"flag"
	"log"
	"os"

	"superservice/internal/superservice"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	app := superservice.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
