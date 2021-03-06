package http

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

const timeout = 60

func TestRepositoryCurrency_GetCurrencies(t *testing.T) {
	t.Parallel()
	t.Run("Should succes when request the currency", func(t *testing.T) {
		t.Parallel()
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			f, err := os.Open("example.json")
			assert.NoError(t, err)
			buffer, err := io.ReadAll(f)
			assert.NoError(t, err)
			if _, err := w.Write(buffer); err != nil {
				w.WriteHeader(http.StatusInternalServerError)

				return
			}
			w.WriteHeader(http.StatusOK)
		}))
		r := NewRepositoryCurrency(server.URL, timeout)
		_, _, _, err := r.GetCurrencies()
		assert.NoError(t, err)
	})
}

func TestNewRepositoryCurrency(t *testing.T) {
	t.Parallel()
	t.Run("Should get new instance of repository", func(t *testing.T) {
		t.Parallel()
		r := NewRepositoryCurrency(faker.URL(), timeout)
		assert.NotNil(t, r.client)
	})
}
