package http

import (
	"boletia/internal"
	"errors"
	"fmt"
	"time"
)

var (
	ErrNotFoundCurrencies = errors.New("error the currencies are not found ")
	ErrNotFoundMeta       = errors.New("error getting meta")
)

type Currency map[string]interface{}

// getCurrency use the map that request return and check the type and make the struct.
func (c Currency) getCurrency() (internal.Currencies, error) {
	currencies := make(internal.Currencies, 0)
	valuesCurrency := c
	for _, currency := range valuesCurrency {
		currencyData := currency.(map[string]interface{})
		curr := internal.Currency{}
		for _, d := range currencyData {
			// Check the type from the value in the map.
			switch d.(type) {
			case string:
				val, ok := d.(string)
				// should be ever true because the switch check the type.
				if !ok {
					return nil, ErrNotFoundCurrencies
				}
				curr.Code = val
			case float64:
				// should be ever true because the switch check the type.
				val, ok := d.(float64)
				if !ok {
					return nil, ErrNotFoundCurrencies
				}
				curr.Value = val
			}
		}
		currencies = append(currencies, curr)
	}

	return currencies, nil
}

// getMeta return the metadata from request.
func (c Currency) getMeta() (*internal.MetaData, error) {
	value, ok := c["last_updated_at"]
	if !ok {
		return nil, ErrNotFoundMeta
	}
	timeUpdateMeta, err := time.Parse(time.RFC3339, value.(string))
	if err != nil {
		return nil, fmt.Errorf("%w:%s", ErrNotFoundMeta, err)
	}

	return &internal.MetaData{
		LastUpdateAt: timeUpdateMeta,
	}, nil
}
