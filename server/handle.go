package server

import (
	"time"

	"golang.org/x/net/websocket"
)

func HandleWs(ws *websocket.Conn) {
	var clientId string
	websocket.Message.Receive(ws, &clientId)
	for !checkLogin(ws, clientId) {
		time.Sleep(1 * time.Second)
	}
	time.Sleep(60 * time.Second)
}
