package internal

import (
	"boletia/kit/event"
	"github.com/google/uuid"
	"time"
)

const CurrencyFailEventType event.Type = "events.currency.fail"

type CurrencyEvent struct {
	event.BaseEvent
	dateExecuteTime time.Time
	duration        time.Duration
	status          string
}

func (e CurrencyEvent) DateExecuteTime() time.Time {
	return e.dateExecuteTime
}

func (e CurrencyEvent) Duration() time.Duration {
	return e.duration
}

func (e CurrencyEvent) Status() string {
	return e.status
}

func NewCurrencyFailEvent(dateExecuteTime time.Time, duration time.Duration, status string) CurrencyEvent {
	return CurrencyEvent{
		BaseEvent:       event.NewBaseEvent(uuid.New().String()),
		dateExecuteTime: dateExecuteTime,
		duration:        duration,
		status:          status,
	}
}

func (e CurrencyEvent) Type() event.Type {
	return CurrencyFailEventType
}
