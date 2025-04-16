package model

import "time"

type Payment struct {
	ID          int       `json:"id"`
	Amount      int       `json:"amount"`
	PaymentDate time.Time `json:"payment_date"`
}
