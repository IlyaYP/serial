package main

import (
	"bytes"
	"fmt"
	"log"

	"go.bug.st/serial"
)

type Port struct {
	port serial.Port

	// Port Name.
	name string
	send chan []byte

	hub *Hub
}

func newPort(hub *Hub) (*Port, error) {
	port, err := serial.Open(hub.port, mode)
	if err != nil {
		return nil, fmt.Errorf("can not open port %s:%w", hub.port, err)
	}

	return &Port{
		hub:  hub,
		port: port,
		name: hub.port,
		send: make(chan []byte, 256),
	}, nil

}

func (p *Port) readPump() {
	defer func() {
		p.port.Close()
	}()

	buff := make([]byte, 100)
	for {
		// Reads up to 100 bytes
		n, err := p.port.Read(buff)
		if err != nil {
			log.Fatal(err)
		}
		if n == 0 {
			fmt.Println("\nEOF")
			break
		}

		fmt.Printf("%s", string(buff[:n]))

		// // If we receive a newline stop reading
		// if strings.Contains(string(buff[:n]), "\n") {
		// 	break
		// }

		buff = bytes.TrimSpace(bytes.Replace(buff, newline, space, -1))
		p.hub.broadcast <- buff

	}
}
