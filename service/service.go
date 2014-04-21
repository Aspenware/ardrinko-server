package service

import (
	"log"
	"net/http"
	"fmt"
	"github.com/AspenWare/ardrinko-server/keg"
)

func Run(statusPipe chan keg.KegStatus) {
	status := keg.KegStatus { }
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Temperature: %f deg, current flow: %d, capacity: %d, available: %d, door: %d\n",
			status.Temperature, status.CurrentFlow, status.Capacity, status.Available, status.Door)
	})
	go monitorStatus(&status, statusPipe)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func monitorStatus(status *keg.KegStatus, statusPipe chan keg.KegStatus) {
	for {
		*status = <-statusPipe
	}
}
