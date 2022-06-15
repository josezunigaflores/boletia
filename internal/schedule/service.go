package schedule

import (
	"boletia/internal"
	"boletia/kit/event"
	"context"
	log "github.com/sirupsen/logrus"
	"time"
)

type Service struct {
	Request    internal.RepositoryHttp
	Repository internal.RepositoryCurrency
	timeOut    int
	eventBus   event.Bus
}

func NewServiceSchedule(request internal.RepositoryHttp, repository internal.RepositoryCurrency, timeOut int, eventBus event.Bus) *Service {
	return &Service{Request: request, Repository: repository, timeOut: timeOut, eventBus: eventBus}
}

func (s Service) Do() {
	background := context.Background()
	for {
		ctx, cancel := context.WithTimeout(background, time.Duration(s.timeOut)*time.Second)
		<-ctx.Done()
		currencies, meta, evnt, err := s.Request.GetCurrencies()
		if err != nil {
			// should save in anyplace
			log.Error(err)
		}
		if err := s.Repository.CreateCurrencies(currencies, *meta); err != nil {
			log.Error(err)
		}
		if err := s.eventBus.Publish(context.Background(), append([]event.Event{}, evnt)); err != nil {
			log.Error(err)
		}
		cancel()
	}
}
