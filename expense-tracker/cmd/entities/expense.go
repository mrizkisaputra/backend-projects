package entities

import "time"

type Expense struct {
	Id          int
	Description string
	Amount      float64
	Date        time.Time
	Category    string
}
