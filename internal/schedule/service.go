package schedule

import (
	"boletia/internal"
	"context"
	"time"
)

type Service struct {
	Request internal.RepositoryHttp
	ctx     context.Context
	timeOut int
}

func (s Service) Do() {
	background := context.Background()
	for {
		ctx, cancel := context.WithTimeout(background, time.Duration(s.timeOut)*time.Second)
		<-ctx.Done()

		cancel()
	}
}
