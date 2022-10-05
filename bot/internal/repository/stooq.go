package repository

import (
	"bot/internal/domain"
	"strings"
)

type StooqRepository interface {
	GetQuote(symbol string) *domain.Quote
}
type stooqRepository struct {
}

func NewStooqRepository() StooqRepository {
	return &stooqRepository{}
}

func (r *stooqRepository) GetQuote(symbol string) *domain.Quote {
	quote := &domain.Quote{}

	quote.Symbol = strings.ToUpper(symbol)
	quote.Close = 123.45

	return quote
}
