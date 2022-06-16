package currency

import (
	"boletia/internal"
)

type Service struct {
	repository internal.RepositoryCurrencyFind
}

func NewServiceCurrency(repository internal.RepositoryCurrencyFind) Service {
	return Service{repository: repository}
}

func (s Service) FindCurrency(code string, finit, fend string) (internal.Currencies, error) {
	c, err := internal.NewCode(code)
	if err != nil {
		return s.repository.Find()
	}
	NewFinit, errFinit := internal.NewTimeFilter(finit)
	NewFend, errFend := internal.NewTimeFilter(fend)

	if c.IsAll() && errFinit != nil && errFend != nil {
		return s.repository.Find()
	}

	if errFinit != nil {
		return nil, errFinit
	}
	if errFend != nil {
		return nil, errFend
	}

	return s.repository.FindWithDate(c, NewFinit, NewFend)
}
