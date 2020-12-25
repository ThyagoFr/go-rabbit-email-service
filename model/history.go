package model

import "time"

type HistoryEmail struct {
	ReceivedAt  time.Time
	UpdatedAt   time.Time
	ServiceName string
	Retry       int
	Status      string
	To          string
	From        string
	Subject     string
}
