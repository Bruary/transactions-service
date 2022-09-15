package models

import "time"

type Transaction struct {
	Amount   float32
	Currency string
	Created  time.Time
	Updated  time.Time
	Deleted  time.Time
}
