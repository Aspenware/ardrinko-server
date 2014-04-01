package keg

import (
	"time"
	"net"
	"bytes"
	"encoding/binary"
)

type KegStatus struct {
	Temperature float32
	CurrentFlow int32
	Capacity uint32
	Available uint32
	LastUpdate time.Time
	Connection* net.UDPConn
}

func Initialize() (KegStatus, error) {
	status := KegStatus {}
	socket, err := net.ListenUDP("udp4", &net.UDPAddr {
		IP: net.ParseIP("0.0.0.0"),
		Port: 59312,
	})
	if err != nil {
		return status, err
	}
	status.Connection = socket
	return status, nil
}

func Monitor(status *KegStatus, eventPipe chan int) {
	buffer := make([]byte, 512)
	for {
		length, _, err := status.Connection.ReadFromUDP(buffer[:])
		if err != nil {
			return // Drop packet
		}
		reader := bytes.NewReader(buffer[:length])
		var temp int32
		err = binary.Read(reader, binary.LittleEndian, &temp)
		status.Temperature = float32(temp) / 100
		if err != nil {
			return // Drop packet
		}
		err = binary.Read(reader, binary.LittleEndian, &status.CurrentFlow)
		if err != nil {
			return // Drop packet
		}
		err = binary.Read(reader, binary.LittleEndian, &status.Capacity)
		if err != nil {
			return // Drop packet
		}
		err = binary.Read(reader, binary.LittleEndian, &status.Available)
		if err != nil {
			return // Drop packet
		}
		status.LastUpdate = time.Now()
		eventPipe <- 1
	}
}
