package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
)

func Echo(ws *websocket.Conn) {
	var err error

	for {
		var reply string

		//这里接受消息 -> websocket.Message.Receive(ws, &reply)
		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}

		fmt.Println("Received back from client: " + reply)
		msg := "Received:  " + reply
		fmt.Println("Sending to client: " + msg)

		websocket.Message.Send(ws, msg)
	}
}

func main() {
	http.Handle("/", websocket.Handler(Echo))

	if err := http.ListenAndServe(":7777", nil); err != nil {
		log.Fatal("listen and server: ", err)
	}
}
