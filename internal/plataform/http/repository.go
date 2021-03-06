package http

import (
	"boletia/internal"
	"boletia/kit/event"
	"encoding/json"
	"io"
	"net"
	"net/http"
	"time"
)

const (
	fail    = "FAIL"
	success = "SUCCESS"
)

type RepositoryCurrency struct {
	path   string
	client *http.Client
}

func NewRepositoryCurrency(path string, timeOut int) RepositoryCurrency {
	c := &http.Client{
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout: time.Duration(timeOut) * time.Second,
			}).Dial,
			TLSHandshakeTimeout: time.Duration(timeOut) * time.Second,
		},
	}
	c.Timeout = time.Duration(timeOut) * time.Second

	return RepositoryCurrency{path: path, client: c}
}

func (rc RepositoryCurrency) GetCurrencies() (internal.Currencies, *internal.MetaData, event.Event, error) {
	resp, err := rc.client.Get(rc.path) // nolint:noctx
	if err != nil {
		return nil, nil, internal.NewCurrencyFailEvent(time.Now().UTC(), rc.client.Timeout, fail), err
	}
	defer resp.Body.Close()
	bts, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, internal.NewCurrencyFailEvent(time.Now().UTC(), rc.client.Timeout, fail), err
	}
	var c struct {
		Meta Currency `json:"meta"`
		Data Currency `json:"data"`
	}
	err = json.Unmarshal(bts, &c)
	if err != nil {
		return nil, nil, internal.NewCurrencyFailEvent(time.Now().UTC(), rc.client.Timeout, fail), err
	}
	currencies, err := c.Data.getCurrency()
	if err != nil {
		return nil, nil, internal.NewCurrencyFailEvent(time.Now().UTC(), rc.client.Timeout, fail), err
	}
	meta, err := c.Meta.getMeta()
	if err != nil {
		return nil, nil, internal.NewCurrencyFailEvent(time.Now().UTC(), rc.client.Timeout, fail), err
	}

	return currencies, meta, internal.NewCurrencyFailEvent(time.Now().UTC(), rc.client.Timeout, success), nil
}
