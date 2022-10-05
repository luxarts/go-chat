package wsserver

import (
	"github.com/gorilla/websocket"
	"time"
)

const (
	writeMaxTime = 5 * time.Second
)

type connection struct {
	wsc  *websocket.Conn
	send chan []byte
}

func (c *connection) write(msgType int, message []byte) error {
	if err := c.wsc.SetWriteDeadline(time.Now().Add(writeMaxTime)); err != nil {
		return err
	}
	return c.wsc.WriteMessage(msgType, message)
}
