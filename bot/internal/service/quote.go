package service

import (
	"bot/internal/domain"
	"bot/internal/repository"
)

type QuoteService interface {
	GetQuote(symbol string) (*domain.Quote, error)
}

type quoteService struct {
	repo repository.StooqRepository
}

func NewQuoteRepository(repo repository.StooqRepository) QuoteService {
	return &quoteService{repo: repo}
}

func (s *quoteService) GetQuote(symbol string) (*domain.Quote, error) {
	return s.repo.GetQuote(symbol)
}
