package main

import (
	"embed"
	"fmt"
	"gochat/db"
	"gochat/server"
	"gochat/utils"
	"io/fs"
	"net/http"
	"os"

	"golang.org/x/net/websocket"
	"gopkg.in/ini.v1"
)

var (
	//go:embed web/*
	webFS embed.FS
)

func syncHandler(ws *websocket.Conn) {
	defer ws.Close()
	utils.Log("handling new websocket connection")
	server.HandleWs(ws)

}

func handleWs(w http.ResponseWriter, r *http.Request) {
	utils.Log(fmt.Sprintf("connection from %s\n", r.RemoteAddr))
	wsHandler := websocket.Handler(syncHandler)
	wsHandler.ServeHTTP(w, r)
	utils.Log("ws connection closed")
}

func main() {
	utils.Log("Ignition")
	if len(os.Args) < 2 {
		fmt.Println("Usage: go&chat <config file>")
		os.Exit(2)
	}
	cfg, err := ini.Load(os.Args[1])
	if err != nil {
		fmt.Printf("Failed to read file: %v\n", err)
		os.Exit(1)
	}
	db.Initialize(cfg)
	defer db.SyncStorage()

	serverRoot, err := fs.Sub(webFS, "web")
	http.Handle("/", http.FileServer(http.FS(serverRoot)))
	http.HandleFunc("/ws", handleWs)
	serverAddress := cfg.Section("server").Key("address").String()
	utils.Log("Listening on " + serverAddress)
	err = http.ListenAndServe(serverAddress, nil)
	if err != nil {
		utils.Fatal(err)
	}
	utils.Log("all done")
}
