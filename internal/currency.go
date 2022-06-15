package internal

import "time"

type Currency struct {
	Code  string
	Value float64
}

type RepositoryHttp interface {
	GetCurrencies() (Currencies, *MetaData, error)
}

type Currencies []Currency

type MetaData struct {
	LastUpdateAt time.Time
}
