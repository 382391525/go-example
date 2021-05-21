package lwebsocket

import "github.com/gorilla/websocket"

type connection struct {
	ws *websocket.Conn

}