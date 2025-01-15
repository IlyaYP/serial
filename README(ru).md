# Сервер WebSocket для последовательной связи

Этот проект представляет собой сервер WebSocket, который взаимодействует с последовательными портами с использованием языка программирования Go. Он использует пакет Gorilla WebSocket для WebSocket-коммуникации и пакет `go.bug.st/serial` для последовательной связи.

## Возможности

- Отображает доступные последовательные порты на сервере.
- Устанавливает WebSocket-соединения для связи с клиентами.
- Настраиваемые параметры последовательной связи.

## Требования

- Go 1.13 или новее
- Пакет Gorilla WebSocket
- Пакет `go.bug.st/serial`

## Установка

1. Клонируйте репозиторий:
    ```sh
    git clone https://github.com/IlyaYP/serial.git
    cd serial
    ```

2. Установите зависимости:
    ```sh
    go get github.com/gorilla/websocket
    go get go.bug.st/serial
    ```

## Использование

1. Соберите проект:
    ```sh
    go build -o serial-server
    ```

2. Запустите сервер:
    ```sh
    ./serial-server -addr=":8080"
    ```

3. Откройте веб-браузер и перейдите по адресу `http://localhost:8080`, чтобы увидеть доступные последовательные порты.

## Конфигурация

Параметры последовательной связи определены в файле `main.go`:
```go
var mode = &serial.Mode{
    BaudRate: 9600,
    Parity:   serial.EvenParity,
    DataBits: 7,
    StopBits: serial.OneStopBit,
}
```

Вы можете изменить эти параметры по мере необходимости.

## Лицензия

Этот проект лицензирован в соответствии с лицензией BSD, которая находится в файле LICENSE.

## Благодарности

- [Gorilla WebSocket](https://github.com/gorilla/websocket)
- [go.bug.st/serial](https://github.com/bugst/go-serial)