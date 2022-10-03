package domain

import "time"

type Data struct {
	User string    `json:"user"`
	Msg  string    `json:"msg"`
	Time time.Time `json:"time"`
}

func (d *Data) IsValid() bool {
	return d.User != "" && d.Msg != ""
}
