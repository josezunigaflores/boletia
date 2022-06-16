package internal

import (
	"boletia/kit/event"
	"errors"
	"fmt"
	"time"
)

type Currency struct {
	Code          string
	Value         float64
	LastUpdatedAt time.Time
}

type RepositoryHttp interface {
	GetCurrencies() (Currencies, *MetaData, event.Event, error)
}

type RepositoryCurrency interface {
	CreateCurrencies(currencies Currencies, data MetaData) error
}

type RepositoryCurrencyFind interface {
	FindWithDate(code Code, finit, fend FilterTime) (Currencies, error)
	Find() (Currencies, error)
}

type Currencies []Currency

type MetaData struct {
	LastUpdateAt time.Time
}

type (
	Code       string
	FilterTime time.Time
)

var (
	ErrBadCode       = errors.New("should not be empty")
	ErrBadTimeFilter = errors.New("the time format is incorrect")
)

func NewCode(code string) (Code, error) {
	if len(code) == 0 {
		return "", ErrBadCode
	}

	return Code(code), nil
}

func (c Code) IsAll() bool {
	return c == "all"
}

const layout = "YYYY-MM-DDThh:mm:ss"

func NewTimeFilter(t string) (FilterTime, error) {
	if len(t) == 0 {
		return FilterTime(time.Now().UTC()), ErrBadTimeFilter
	}
	newTime, err := time.Parse(layout, t)
	if err != nil {
		return FilterTime(newTime), fmt.Errorf("%w:%s", ErrBadTimeFilter, err)
	}

	return FilterTime(newTime), nil
}
