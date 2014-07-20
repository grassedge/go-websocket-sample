package chat

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
)

type Client struct {
	Id int
	ws *websocket.Conn
	// share with server
	removeClientCh chan *Client
	messageCh      chan string
}

func NewClient(ws *websocket.Conn, remove chan *Client, message chan string) *Client {
	return &Client{
		ws:             ws,
		removeClientCh: remove,
		messageCh:      message,
	}
}

func (client *Client) Start() {
	for {
		var message string
		err := websocket.Message.Receive(client.ws, &message)
		if err != nil {
			client.removeClientCh <- client
			return
		} else {
			client.messageCh <- message
		}
	}
}

func (client *Client) Send(message string) {
	err := websocket.Message.Send(client.ws, message)
	if err != nil { fmt.Println(message) }
}

func (client *Client) Close() {
	client.ws.Close()
}
