package currency

import (
	"boletia/internal/utils"
	"boletia/kit/command/commandmocks"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var errAnythingTest = errors.New("error anything")

func TestGetCurrencies(t *testing.T) {
	t.Run("should get all currenties", func(t *testing.T) {
		w := httptest.NewRecorder()
		_, engine := gin.CreateTestContext(w)
		any := &commandmocks.Bus{}
		any.On("Dispatch", mock.Anything, mock.Anything).Return(utils.NewBadRequest(errAnythingTest), nil)
		engine.GET("", GetCurrencies(any))
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		engine.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
