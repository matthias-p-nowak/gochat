package server

import (
	"context"
	"gochat/db"
	"gochat/db/ent/usertoken"

	"golang.org/x/net/websocket"
)

func checkLogin(ws *websocket.Conn, clientId string) bool {
	ctx := context.Background()
	_, err := db.DbStorage.UserToken.Query().Where(usertoken.Token(clientId)).First(ctx)
	if err != nil {
		// TODO: store user information
		return true
	}
	return false
}
