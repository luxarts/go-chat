package domain

import "time"

type Data struct {
	User string    `json:"user"`
	Msg  string    `json:"msg"`
	Time time.Time `json:"time"`
}
