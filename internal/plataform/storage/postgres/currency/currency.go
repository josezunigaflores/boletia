package currency

import "time"

const sqlTableName = "currency"

type SQLCurrency struct {
	Code          string
	Value         float64
	LastUpdatedAt time.Time
}
