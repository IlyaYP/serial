package main

import (
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

// newPort creates a new port and returns a pointer to it.
// It returns an error if the port can not be opened.
func newPort(hub *Hub) (*Port, error) {
	port, err := serial.Open(hub.portName, mode)
	if err != nil {
		return nil, fmt.Errorf("can not open port %s:%w", hub.portName, err)
	}

	return &Port{
		hub:  hub,
		port: port,
		name: hub.portName,
		send: make(chan []byte, 256),
	}, nil

}

// writePump sends messages to the port.
// It closes the port when the channel is closed.
// It sends messages to the port until the channel is closed.
func (p *Port) writePump() {
	defer func() {
		p.port.Close()
	}()
	for {
		select {
		case message, ok := <-p.send:
			if !ok {
				// The hub closed the channel.
				return
			}
			// message = append(message, "\n\r"...)
			n, err := p.port.Write(message)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Sent %v bytes\n", n)
		}
	}

}

// readPump reads messages from the port.
// It sends messages to the hub.
// It closes the port when the channel is closed.
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
		// fmt.Printf("%s", string(buff[:n]))

		c := make([]byte, n)
		copy(c, buff[:n])

		p.hub.broadcast <- c

	}
}
