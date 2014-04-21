package service

import (
	"log"
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/AspenWare/ardrinko-server/keg"
)

type StatusJson struct {
	Temp float32
	Flow float32
	Capacity float32
	Available float32
	Door bool
}

func Run(statusPipe chan keg.KegStatus) {
	status := keg.KegStatus { }
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("ContentType", "application/json")
		data := StatusJson { status.Temperature, status.CurrentFlow, status.Capacity, status.Available, status.Door == 1 }
		j, err := json.Marshal(data)
		if err != nil {
			fmt.Fprintf(w, "{\"error\":true}")
		} else {
			fmt.Fprintf(w, string(j))
		}
	})
	go monitorStatus(&status, statusPipe)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func monitorStatus(status *keg.KegStatus, statusPipe chan keg.KegStatus) {
	for {
		*status = <-statusPipe
	}
}
