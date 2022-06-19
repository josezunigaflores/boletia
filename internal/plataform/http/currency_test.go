package http

import (
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestCurrency_getCurrency(t *testing.T) {
	t.Parallel()
	t.Run("should get changes from map", func(t *testing.T) {
		t.Parallel()
		c := make(Currency)
		data := make(map[string]interface{})
		data["code"] = "MX"
		data["value"] = 190.00
		c["MX"] = data

		value, err := c.getCurrency()
		assert.NoError(t, err)
		log.Info(value)
	})
}

func TestCurrency_getMeta(t *testing.T) {
	t.Parallel()
	t.Run("should get meta data", func(t *testing.T) {
		t.Parallel()
		c := make(Currency)
		c["last_updated_at"] = "2022-06-13T23:59:59Z"
		m, err := c.getMeta()
		assert.NoError(t, err)
		log.Info(m)
	})
	t.Run("should get error because the format of the time is incorrect", func(t *testing.T) {
		t.Parallel()
		c := make(Currency)
		c["last_updated_at"] = time.Now().UTC().String()
		m, err := c.getMeta()
		assert.Error(t, err)
		log.Info(m)
	})

	t.Run("should get error because the field don't exist in the map", func(t *testing.T) {
		t.Parallel()
		c := make(Currency)
		c["any"] = time.Now().UTC().String()
		m, err := c.getMeta()
		assert.Error(t, err)
		log.Info(m)
	})
}
