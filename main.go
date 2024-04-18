package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"slices"

	"go.bug.st/serial"
)

var PORTS []string

var mode = &serial.Mode{
	BaudRate: 115200,
	Parity:   serial.EvenParity,
	DataBits: 7,
	StopBits: serial.OneStopBit,
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "method not found", http.StatusNotFound)
		return
	}

	log.Printf("port: %v", r.URL.Query().Get("port"))

	if !slices.Contains(PORTS, r.URL.Query().Get("port")) {
		// http.ServeFile(w, r, "templates/index.html")
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			log.Fatalf("template parsing: %s", err)
		}

		err = tmpl.Execute(w, PORTS)
		if err != nil {
			log.Fatalf("template execution: %s", err)
		}
	}

	log.Printf("url: %v", r.URL.String())

}

func main() {
	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		log.Fatal("No serial ports found!")
	}
	for _, port := range ports {
		fmt.Printf("Found port: %v\n", port)
	}
	PORTS = ports

	hub := NewHub()
	go hub.run()

	http.HandleFunc("/", serveIndex)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})

	log.Fatal(http.ListenAndServe(":3001", nil))
}
