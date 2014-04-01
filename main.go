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
	eventPipe := make(chan int)
	go keg.Monitor(&status, eventPipe)
	log.Println("Listening for keg on " + status.Connection.LocalAddr().String())
	for {
		<-eventPipe
		log.Printf("Temperature: %f deg, current flow: %d, capacity: %d, available: %d\n",
			status.Temperature, status.CurrentFlow, status.Capacity, status.Available)
	}
}
