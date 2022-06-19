package currency

import (
	"boletia/internal/plataform/storage/postgres"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func TestCurrencies_ToCurrencies(t *testing.T) {
	t.Parallel()
	t.Run("should became to internal currencies", func(t *testing.T) {
		t.Parallel()
		c := make(Currencies, 0)
		want := SQLCurrency{
			Model: postgres.Model{},
			Code:  faker.Word(),
			Value: 0,
		}
		c = append(c, want)
		newC, err := c.ToCurrencies()
		assert.NoError(t, err)
		got := newC[0]
		if assert.NotNil(t, got) {
			return
		}
		assert.Equal(t, want.Code, got.Code)
		assert.Equal(t, want.Value, got.Value)
	})
}
