package wsserver

import (
	"backend/internal/domain"
	"encoding/json"
	"errors"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

const (
	pongTime   = 60 * time.Second
	pingTime   = 55 * time.Second
	maxMsgSize = 512
)

type subscription struct {
	conn *connection
}

func (s *subscription) writeRoutine() {
	go func() {
		c := s.conn
		ticker := time.NewTimer(pingTime)
		defer func() {
			ticker.Stop()
			if err := c.wsc.Close(); err != nil {
				log.Printf("error closing wsc: %v\n", err)
			}
		}()

		for {
			select {
			case m, ok := <-c.send:
				if !ok {
					if err := c.write(websocket.CloseMessage, []byte{}); err != nil {
						log.Printf("error sending close message: %v\n", err)
					}
					return
				}
				if err := c.write(websocket.TextMessage, m); err != nil {
					log.Printf("error sending text message: %v\n", err)
					return
				}
			case <-ticker.C:
				if err := c.write(websocket.PingMessage, []byte{}); err != nil {
					log.Printf("error sending ping message: %v\n", err)
					return
				}
			}
		}
	}()
}
func (s *subscription) readRoutine(unregisterChan chan subscription, dataChan chan domain.Data) {
	go func() {
		c := s.conn
		defer func() {
			unregisterChan <- *s
			if err := c.wsc.Close(); err != nil {
				log.Printf("error closing wsc: %v\n", err)
			}
		}()

		c.wsc.SetReadLimit(maxMsgSize)
		if err := c.wsc.SetReadDeadline(time.Now().Add(pongTime)); err != nil {
			log.Printf("error setting read deadline: %v\n", err)
		}
		c.wsc.SetPongHandler(func(string) error {
			return c.wsc.SetReadDeadline(time.Now().Add(pongTime))
		})

		for {
			msgType, dataBytes, err := c.wsc.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
					log.Printf("unexpected error on close: %v\n", err)
				}
				break
			}

			if msgType == websocket.TextMessage {
				if err := s.messageHandler(dataBytes, dataChan); err != nil {
					log.Printf("error handling msg: %v\n", err)
				}
			}
		}
	}()
}

func (s *subscription) messageHandler(dataBytes []byte, dataChan chan domain.Data) error {
	var data domain.Data

	if err := json.Unmarshal(dataBytes, &data); err != nil {
		log.Printf("error unmarhaling: %v\n", err)
		return err
	}

	if !data.IsValid() {
		return errors.New("empty user or msg")
	}

	data.Time = time.Now().UTC()

	dataChan <- data

	return nil
}
