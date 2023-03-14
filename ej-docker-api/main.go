package main

import (
	"github.com/emikohmann/ucc-arqsoft-1/ej-docker-api/app"
	"log"
)

func main() {
	if err := app.RunApp(); err != nil {
		log.Fatal(err)
	}
}
