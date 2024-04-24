// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"sync"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	sync.RWMutex

	port *Port

	// Port Name.
	portName string

	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func newHub(portName string) (*Hub, error) {

	hub := &Hub{
		portName:   portName,
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    map[*Client]bool{},
	}

	Port, err := newPort(hub)
	if err != nil {
		return nil, fmt.Errorf("can not create hub for %s:%w", portName, err)
	}

	hub.port = Port

	return hub, nil
}

func (h *Hub) run() {
	go h.port.readPump()
	go h.port.writePump()

	for {
		select {
		case client := <-h.register:
			h.Lock()
			h.clients[client] = true
			h.Unlock()
			log.Printf("client registered %s %s", h.portName, client.id)

		case client := <-h.unregister:
			h.Lock()
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
				log.Printf("client unregistered %s", client.id)
			}
			h.Unlock()
		case message := <-h.broadcast:
			// fmt.Printf("%s", string(message))
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
