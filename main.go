package main

import (
	"./chat"
	"net/http"
)

func main() {
	server := chat.NewServer()
	go server.Start()

	http.Handle("/ws", server.WebsocketHandler())
	http.Handle("/", http.FileServer(http.Dir("webroot")))
	http.ListenAndServe(":8248", nil)
}
