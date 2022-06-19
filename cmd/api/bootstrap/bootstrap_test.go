package bootstrap

import (
	"boletia/internal"
	currency2 "boletia/internal/currency"
	"boletia/internal/mocks"
	"boletia/internal/plataform/bus/inmemory"
	"boletia/internal/plataform/server/handler/currency"
	"github.com/bxcodec/faker/v3"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetCurrenciesWithoutDBIIntegration(t *testing.T) {
	t.Parallel()
	t.Run("should use all implements with command and events", func(t *testing.T) {
		t.Parallel()
		bus := inmemory.NewCommandBus()
		find := &mocks.RepositoryCurrencyFind{}
		currencies := make(internal.Currencies, 0)
		currencies = append(currencies, internal.Currency{
			Code:          faker.Word(),
			Value:         0,
			LastUpdatedAt: time.Now().UTC(),
		})
		find.On("FindWithDate", mock.Anything, mock.Anything, mock.Anything).Return(currencies, nil)
		serviceCurrency := currency2.NewServiceCurrency(find)
		cmd := currency2.NewCurrencyHandler(serviceCurrency)
		bus.Register(currency2.CurrencyCommandType, cmd)
		w := httptest.NewRecorder()
		_, engine := gin.CreateTestContext(w)
		engine.GET("currencies/:id", currency.GetCurrencies(bus))

		req := httptest.NewRequest(http.MethodGet, "/currencies/mxn", nil)
		q := req.URL.Query()
		q.Add("finit", "2021-10-28T18:15:00")
		q.Add("fend", "2021-10-28T20:15:00")
		req.URL.RawQuery = q.Encode()
		engine.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}
