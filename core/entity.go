package core

import (
	"net/mail"
	"time"
)

type (
	Mail struct {
		From    *mail.Address
		To      []*mail.Address
		CC      []*mail.Address
		Time    time.Time
		Subject string
		Text    string
	}
)
