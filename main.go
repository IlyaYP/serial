// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"go.bug.st/serial"
)

var PORTS []string

var mode = &serial.Mode{
	BaudRate: 115200,
	Parity:   serial.EvenParity,
	DataBits: 7,
	StopBits: serial.OneStopBit,
}

var addr = flag.String("addr", ":8080", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
}

func main() {
	flag.Parse()

	ports := []string{"COM1", "COM2", "COM3"} //tmp
	// ports, err := serial.GetPortsList()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	if len(ports) == 0 {
		log.Fatal("No serial ports found!")
	}

	hubs := make(map[string]*Hub)
	for _, port := range ports {
		log.Printf("Found port: %v\n", port)
		hubs[port] = newHub(port)
		go hubs[port].run()
	}
	PORTS = ports

	// hub := newHub()
	// go hub.run()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hubs, w, r)
	})
	server := &http.Server{
		Addr:              *addr,
		ReadHeaderTimeout: 3 * time.Second,
	}
	log.Fatal(server.ListenAndServe())

}
