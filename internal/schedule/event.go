package schedule

import (
	"boletia/internal"
	"boletia/kit/event"
	"context"
	"errors"
)

type Event struct {
	repository internal.RepositoryCalls
}

func NewEvent(repository internal.RepositoryCalls) *Event {
	return &Event{repository: repository}
}

func (e Event) Handle(_ context.Context, evt event.Event) error {
	courseCreatedEvt, ok := evt.(internal.CurrencyEvent)
	if !ok {
		return errors.New("unexpected event")
	}
	return e.repository.CreateCall(courseCreatedEvt.DateExecuteTime(), courseCreatedEvt.Duration(), courseCreatedEvt.Status())
}
