package domain

import "time"

type Quote struct {
	Symbol string
	Date   time.Time
	Open   float32
	High   float32
	Low    float32
	Close  float32
	Volume uint64
}
