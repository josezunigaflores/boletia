package currency

import (
	"boletia/internal"
	"boletia/internal/plataform/storage/postgres"
	"time"
)

const sqlTableName = "currency"

type SQLCurrency struct {
	postgres.Model
	Code          string
	Value         float64
	LastUpdatedAt time.Time
}

type Currencies []SQLCurrency

func (c Currencies) ToCurrencies() (internal.Currencies, error) {
	currencies := make(internal.Currencies, 0)
	for _, currency := range c {
		currencies = append(currencies, internal.Currency{
			Code:          currency.Code,
			Value:         currency.Value,
			LastUpdatedAt: currency.LastUpdatedAt,
		})
	}

	return currencies, nil
}
