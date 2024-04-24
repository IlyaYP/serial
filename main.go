// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"time"

	"go.bug.st/serial"
)

var PORTS []string

var mode = &serial.Mode{
	BaudRate: 9600,
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
	// http.ServeFile(w, r, "home.html")
	tmpl, err := template.ParseFiles("home.html")
	if err != nil {
		log.Fatalf("template parsing: %s", err)
	}

	err = tmpl.Execute(w, PORTS)
	if err != nil {
		log.Fatalf("template execution: %s", err)
	}

}

func main() {
	flag.Parse()

	// ports := []string{"COM1", "COM2", "COM3"} //tmp
	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		log.Fatal("No serial ports found!")
	}

	hubs := make(map[string]*Hub)
	for _, port := range ports {
		log.Printf("Found port: %v\n", port)
		hub, err := newHub(port)
		if err != nil {
			log.Printf("%s\n", err)
			continue
		}
		hubs[port] = hub
		PORTS = append(PORTS, port)
		go hub.run()
	}

	if len(PORTS) == 0 {
		log.Fatal("No free ports found!")
	}

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
