package calls

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestRepository_CreateCall(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	require.NoError(t, err)

	s := NewCallRepository(gormDB)
	var (
		dateExecuteTime = time.Now()
		duration        = time.Duration(10)
		status          = faker.Word()
	)

	assert.NoError(t, err)
	q := `INSERT INTO "call" ("date_execute_time","duration","status") VALUES ($1,$2,$3)`
	mock.ExpectBegin()
	mock.ExpectExec(q).
		WithArgs(dateExecuteTime, duration, status).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	err = s.CreateCall(dateExecuteTime, duration, status)
	assert.NoError(t, mock.ExpectationsWereMet())
	assert.NoError(t, err)
}
