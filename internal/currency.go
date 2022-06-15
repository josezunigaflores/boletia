package internal

import (
	"boletia/kit/event"
	"time"
)

type Currency struct {
	Code  string
	Value float64
}

type RepositoryHttp interface {
	GetCurrencies() (Currencies, *MetaData, event.Event, error)
}

type RepositoryCurrency interface {
	CreateCurrencies(currencies Currencies, data MetaData) error
}

type Currencies []Currency

type MetaData struct {
	LastUpdateAt time.Time
}
