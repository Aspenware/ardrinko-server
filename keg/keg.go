package keg

import (
	"time"
	"net"
	"bytes"
	"encoding/binary"
	. "github.com/AspenWare/ardrinko-server/config"
)

type KegStatus struct {
	Temperature float32
	CurrentFlow float32
	Capacity float32
	Available float32
	Door uint32
	LastUpdate time.Time
	Connection* net.UDPConn
}

func Initialize(config *Config) (KegStatus, error) {
	status := KegStatus {}
	socket, err := net.ListenUDP("udp4", &net.UDPAddr {
		IP: net.ParseIP(config.UDP.Endpoint),
		Port: config.UDP.Port,
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
		var flow int32
		var capacity int32
		var available int32
		err = binary.Read(reader, binary.LittleEndian, &temp)
		status.Temperature = float32(temp) / 100
		if err != nil {
			continue // Drop packet
		}
		err = binary.Read(reader, binary.LittleEndian, &flow)
		status.CurrentFlow = float32(flow) / 100
		if err != nil {
			continue // Drop packet
		}
		err = binary.Read(reader, binary.LittleEndian, &capacity)
		status.Capacity = float32(capacity) / 100
		if err != nil {
			continue // Drop packet
		}
		err = binary.Read(reader, binary.LittleEndian, &available)
		status.Available = float32(available) / 100
		if err != nil {
			continue // Drop packet
		}
		err = binary.Read(reader, binary.LittleEndian, &status.Door)
		if err != nil {
			continue // Drop packet
		}
		status.LastUpdate = time.Now()
		eventPipe <- 1
	}
}
