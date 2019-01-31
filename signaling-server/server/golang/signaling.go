package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"log"
	"net/http"
	"sync"
)

var clients map[*websocket.Conn]struct{} = make(map[*websocket.Conn]struct{})
var mutex sync.RWMutex

func notify(sender *websocket.Conn, message string) error {
	mutex.RLock()

	for c := range clients {
		if c != sender {
			if err := websocket.Message.Send(c, message); err != nil {
				return err
			}
		}
	}
	mutex.RUnlock()

	return nil
}

func handler(conn *websocket.Conn) {
	log.Println("-- websocket connected --")

	mutex.Lock()
	clients[conn] = struct{}{}
	mutex.Unlock()

	defer func(ws *websocket.Conn) {
		ws.Close()
		mutex.Lock()
		delete(clients, ws)
		mutex.Unlock()
	}(conn)

	var message string
	for {
		if err := websocket.Message.Receive(conn, &message); err != nil {
			if err == io.EOF {
				log.Println("- disconnected -")
			} else {
				log.Println(fmt.Sprintf("- unexpected error: %v -", err))
			}
			return
		}

		log.Println(fmt.Sprintf("- %s -", message))
		if err := notify(conn, message); err != nil {
			log.Println("notify error: ", err)
			return
		}
	}
}

func main() {
	s := websocket.Server{Handler: handler}
	http.Handle("/", s)
	err := http.ListenAndServe(":3001", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}

	log.Println("websocket server start. port=3001")
}
