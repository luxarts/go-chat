package repository

import (
	"backend/internal/domain"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMessagesRepository_Add_ReadAll(t *testing.T) {
	// Given
	msgs := []domain.Data{
		{User: "1", Msg: "a"},
		{User: "2", Msg: "b"},
		{User: "3", Msg: "c"},
		{User: "4", Msg: "d"},
	}

	r := NewMessagesRepository()

	// When
	r.Add(&msgs[0])
	r.Add(&msgs[1])
	r.Add(&msgs[2])
	r.Add(&msgs[3])

	msgsRead := r.ReadAll()

	// Then
	require.Equal(t, maxStoredMessages, len(msgsRead))
	require.Equal(t, msgs[1], msgsRead[0])
	require.Equal(t, msgs[2], msgsRead[1])
	require.Equal(t, msgs[3], msgsRead[2])
}
