# Serial Communication WebSocket Server

This project is a WebSocket server that communicates with serial ports using the Go programming language. It uses the Gorilla WebSocket package for WebSocket communication and the `go.bug.st/serial` package for serial communication.

## Features

- Lists available serial ports on the server.
- Establishes WebSocket connections to communicate with clients.
- Configurable serial communication parameters.

## Requirements

- Go 1.13 or later
- Gorilla WebSocket package
- `go.bug.st/serial` package

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/IlyaYP/serial.git
    cd serial
    ```

2. Install the dependencies:
    ```sh
    go get github.com/gorilla/websocket
    go get go.bug.st/serial
    ```

## Usage

1. Build the project:
    ```sh
    go build -o serial-server
    ```

2. Run the server:
    ```sh
    ./serial-server -addr=":8080"
    ```

3. Open a web browser and navigate to `http://localhost:8080` to see the available serial ports.

## Configuration

The serial communication parameters are defined in the `main.go` file:
```go
var mode = &serial.Mode{
    BaudRate: 9600,
    Parity:   serial.EvenParity,
    DataBits: 7,
    StopBits: serial.OneStopBit,
}
```

You can modify these parameters as needed.

## License

This project is licensed under the BSD-style license found in the LICENSE file.

## Acknowledgements

- [Gorilla WebSocket](https://github.com/gorilla/websocket)
- [go.bug.st/serial](https://github.com/bugst/go-serial)