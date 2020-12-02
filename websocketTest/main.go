package main

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)
func wsHandler(w http.ResponseWriter, r *http.Request)  {
	var (
		err error
		conn *websocket.Conn
		data []byte
	)
	// upgrade: websocket
	if conn, err = upgrader.Upgrade(w, r, nil); err != nil {
		return
	}
	for {
		// 获取到数据
		if _, data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}
		// 发送数据
		if err = conn.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}
	}
	ERR:
		conn.Close()
}
func main() {
	// http://localhost:7777/ws
	http.HandleFunc("/ws", wsHandler)
	http.ListenAndServe(":7777",nil)
}
