package internal

import "time"

type Call struct {
	DateExecuteTime time.Time
	Duration        time.Duration
	Status          string
}

type Calls []Call

//go:generate mockery --case=snake --outpkg=mocks --output=mocks --name=RepositoryCalls
type RepositoryCalls interface {
	CreateCall(dateExecuteTime time.Time, duration time.Duration, status string) error
}
