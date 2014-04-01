package keg

import (
	"time"
	"net"
	"log"
)

type KegStatus struct {
	Temperature float64
	CurrentFlow float64
	Capacity float64
	Available float64
	LastUpdate time.Time
	Connection* net.UDPConn
}

func Initialize() (KegStatus, error) {
	status := KegStatus {}
	socket, err := net.ListenUDP("udp4", &net.UDPAddr {
		IP: net.ParseIP("127.0.0.1"), //net.IPv4bcast,
		Port: 59312,
	})
	if err != nil {
		return status, err
	}
	status.Connection = socket
	return status, nil
}

func Monitor(status KegStatus) {
	var buffer [512]byte
	for {
		length, from, err := status.Connection.ReadFromUDP(buffer[:])
		if err != nil {
			log.Fatal(err)
		}
		data := string(buffer[:length])
		log.Print("<<[" + from.IP.String() + "]: " + data)
	}
}
