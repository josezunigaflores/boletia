package calls

import "time"

const sqlTableName = "call"

type SQLCall struct {
	DateExecuteTime time.Time
	Duration        time.Duration
	Status          string
}
