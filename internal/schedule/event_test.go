package schedule

import (
	"boletia/internal"
	"boletia/internal/mocks"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestEvent_Handle(t *testing.T) {
	t.Parallel()
	t.Run("Should call to create call repository", func(t *testing.T) {
		t.Parallel()
		rc := &mocks.RepositoryCalls{}
		rc.On("CreateCall", mock.Anything, mock.Anything, mock.Anything).
			Return(nil)
		evt := NewEvent(rc)
		err := evt.Handle(context.Background(), internal.CurrencyEvent{})
		assert.NoError(t, err)
	})
	t.Run("Should return error because the event is incorrect", func(t *testing.T) {
		t.Parallel()
		rc := &mocks.RepositoryCalls{}
		rc.On("CreateCall", mock.Anything, mock.Anything, mock.Anything).
			Return(nil)
		evt := NewEvent(rc)
		err := evt.Handle(context.Background(), &internal.CurrencyEvent{})
		assert.Error(t, err)
	})
}
