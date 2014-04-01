package main

import (
	"log"
	"github.com/AspenWare/ardrinko-server/keg"
)

func main() {
	log.Println("Spinning up Ardrinko server...")
	status, err := keg.Initialize()
	if err != nil {
		log.Fatal(err)
	}
	go keg.Monitor(status)
	log.Println("Listening for keg on " + status.Connection.LocalAddr().String())
	select {
	}
}
