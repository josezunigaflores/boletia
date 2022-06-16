package calls

import "time"

const sqlTableName = "sql_call"

type SQLCall struct {
	DateExecuteTime time.Time
	Duration        time.Duration
	Status          string
}
