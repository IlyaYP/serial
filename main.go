package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"go.bug.st/serial"
)

var PORTS []string

func serveIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "method not found", http.StatusNotFound)
		return
	}

	// http.ServeFile(w, r, "templates/index.html")
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatalf("template parsing: %s", err)
	}

	mode := &serial.Mode{
		BaudRate: 115200,
		Parity:   serial.EvenParity,
		DataBits: 7,
		StopBits: serial.OneStopBit,
	}

	err = tmpl.Execute(w, PORTS)
	if err != nil {
		log.Fatalf("template execution: %s", err)
	}

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

	http.HandleFunc("/", serveIndex)
	log.Fatal(http.ListenAndServe(":3001", nil))

}
