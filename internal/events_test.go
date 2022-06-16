package internal

import (
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCurrencyEvent_DateExecuteTime(t *testing.T) {
	t.Parallel()
	t.Run("should return DateExecuteTime field", func(t *testing.T) {
		t.Parallel()
		det := time.Now()
		ce := NewCurrencyFailEvent(det, time.Duration(90), faker.Word())
		assert.Equal(t, det, ce.DateExecuteTime())
		assert.Equal(t, ce.duration, ce.Duration())
		assert.Equal(t, ce.status, ce.Status())
		assert.Equal(t, CurrencyFailEventType, ce.Type())
	})
}
