package main

import (
	"container/list"
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"log"
	"net/http"
)

var clients = list.List{}

func notify(sender *websocket.Conn, message string) {
	for e := clients.Front(); e != nil; e = e.Next() {
		if e.Value != sender {
			websocket.Message.Send(e.Value.(*websocket.Conn), message)
		}
	}
}

func handler(conn *websocket.Conn) {
	log.Println("-- websocket connected --")

	e := clients.PushBack(conn)

	defer func() {
		conn.Close()
		clients.Remove(e)
	}()

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
		notify(conn, message)
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
