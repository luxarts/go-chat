package service

import (
	"backend/internal/repository"
	"backend/internal/wsserver"
	"net/http"
)

type ChatRoomsService interface {
	Connect(r http.ResponseWriter, writer *http.Request)
}

type chatRoomsService struct {
	wss     wsserver.WSServer
	msgRepo repository.MessagesRepository
}

func NewChatRoomsService(wss wsserver.WSServer, msgRepo repository.MessagesRepository) ChatRoomsService {
	return &chatRoomsService{
		wss:     wss,
		msgRepo: msgRepo,
	}
}

func (srv *chatRoomsService) Connect(r http.ResponseWriter, w *http.Request) {
	// Create connection
	conn := srv.wss.CreateConnection(r, w)

	// Subscribe
	srv.wss.Subscribe(conn)

	// Send history
	msgs := srv.msgRepo.ReadAll()

	for _, m := range msgs {
		srv.wss.SendMessage(conn, &m)
	}
}
