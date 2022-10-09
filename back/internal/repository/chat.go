package repository

import (
	"backend/internal/domain"
	"sync"
)

const maxStoredMessages = 50

type MessagesRepository interface {
	Add(d *domain.Data)
	ReadAll() []domain.Data
}

type messagesRepository struct {
	mu         sync.RWMutex
	messagesDB []domain.Data
}

func NewMessagesRepository() MessagesRepository {
	return &messagesRepository{
		messagesDB: make([]domain.Data, 0),
	}
}

func (r *messagesRepository) Add(d *domain.Data) {
	r.mu.Lock()

	r.messagesDB = append(r.messagesDB, *d)

	if len(r.messagesDB) > maxStoredMessages {
		r.messagesDB = r.messagesDB[1:]
	}

	r.mu.Unlock()
}
func (r *messagesRepository) ReadAll() []domain.Data {
	return r.messagesDB
}
