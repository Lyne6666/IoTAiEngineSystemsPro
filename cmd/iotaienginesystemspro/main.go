// cmd/iotaienginesystemspro/main.go
package main

import (
	"flag"
	"log"
	"os"
	
	"iotaienginesystemspro/internal/iotaienginesystemspro"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()
	
	app := iotaienginesystemspro.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
