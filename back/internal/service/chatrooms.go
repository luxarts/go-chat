package service

import (
	"backend/internal/wsserver"
	"net/http"
)

type ChatRoomsService interface {
	Connect(r http.ResponseWriter, writer *http.Request)
}

type chatRoomsService struct {
	wss wsserver.WSServer
}

func NewChatRoomsService(wss wsserver.WSServer) ChatRoomsService {
	return &chatRoomsService{
		wss: wss,
	}
}

func (srv *chatRoomsService) Connect(r http.ResponseWriter, w *http.Request) {
	// Create connection
	conn := srv.wss.CreateConnection(r, w)

	// Subscribe
	srv.wss.Subscribe(conn)
}
