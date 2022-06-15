package http

import (
	"boletia/internal"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type RepositoryCurrency struct {
	path   string
	client http.Client
}

func NewRepositoryCurrency(path string, timeOut int) RepositoryCurrency {
	c := http.Client{}
	c.Timeout = time.Duration(timeOut) * time.Second
	return RepositoryCurrency{path: path, client: c}
}

func (rc RepositoryCurrency) GetCurrencies() (internal.Currencies, *internal.MetaData, error) {
	resp, err := rc.client.Get(rc.path)
	if err != nil {
		return nil, nil, err
	}
	bts, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	var c struct {
		Meta Currency `json:"meta"`
		Data Currency `json:"data"`
	}
	err = json.Unmarshal(bts, &c)
	if err != nil {
		return nil, nil, err
	}
	currencies, err := c.Data.getCurrency()
	if err != nil {
		return nil, nil, err
	}
	meta, err := c.Meta.getMeta()
	if err != nil {
		return nil, nil, err
	}

	return currencies, meta, nil
}
