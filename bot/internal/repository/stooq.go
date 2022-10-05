package repository

import (
	"bot/internal/defines"
	"bot/internal/domain"
	"bytes"
	"encoding/csv"
	"errors"
	"github.com/go-resty/resty/v2"
	"net/url"
	"strconv"
	"time"
)

const (
	paramSymbol       = "s"
	paramHeaders      = "h"
	paramFormat       = "f"
	dateAndTimeFormat = "sd2t2ohlcv"
)

type StooqRepository interface {
	GetQuote(symbol string) (*domain.Quote, error)
}
type stooqRepository struct {
	rc *resty.Client
}

func NewStooqRepository(rc *resty.Client) StooqRepository {
	return &stooqRepository{
		rc: rc,
	}
}

func (r *stooqRepository) GetQuote(symbol string) (*domain.Quote, error) {
	req := r.rc.R()

	req.SetQueryParam(paramSymbol, symbol)
	req.SetQueryParam(paramHeaders, "")
	req.SetQueryParam(paramFormat, dateAndTimeFormat)

	u := url.URL{
		Scheme: "https",
		Host:   defines.APIStooqURL,
		Path:   defines.APIStooqPathGetQuote,
	}

	resp, err := req.Get(u.String())
	if err != nil {
		return nil, err
	}

	return parseCSVToQuote(resp.Body())
}

func parseCSVToQuote(data []byte) (*domain.Quote, error) {
	r := csv.NewReader(bytes.NewBuffer(data))

	// Build a map of the headers with the values to not depend on the order of the CSV values in parsing
	quoteMap := make(map[string]string)
	headers, err := r.Read()
	if err != nil {
		return nil, err
	}

	for _, h := range headers {
		quoteMap[h] = ""
	}

	values, err := r.Read()
	if err != nil {
		return nil, err
	}

	for i, h := range headers {
		if values[i] == "N/D" { // Stooq returns the CSV with N/D values if symbol doesn't exists
			return nil, errors.New("symbol not found")
		}
		quoteMap[h] = values[i]
	}

	quote := domain.Quote{
		Symbol: quoteMap["Symbol"],
	}

	quote.Date, err = time.Parse("2006-01-02 15:04:05", quoteMap["Date"]+" "+quoteMap["Time"])
	if err != nil {
		return nil, err
	}
	quote.Open, err = strconv.ParseFloat(quoteMap["Open"], 64)
	if err != nil {
		return nil, err
	}
	quote.High, err = strconv.ParseFloat(quoteMap["High"], 64)
	if err != nil {
		return nil, err
	}
	quote.Low, err = strconv.ParseFloat(quoteMap["Low"], 64)
	if err != nil {
		return nil, err
	}
	quote.Close, err = strconv.ParseFloat(quoteMap["Close"], 64)
	if err != nil {
		return nil, err
	}
	quote.Volume, err = strconv.ParseInt(quoteMap["Volume"], 10, 64)
	if err != nil {
		return nil, err
	}

	return &quote, nil
}
