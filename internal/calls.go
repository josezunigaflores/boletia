package internal

import "time"

type Call struct {
	DateExecuteTime time.Time
	Duration        time.Duration
	Status          string
}

type Calls []Call
