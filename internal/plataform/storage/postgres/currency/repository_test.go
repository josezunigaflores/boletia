package currency

import (
	"boletia/internal"
	"database/sql/driver"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestRepository_CreateCurrencies(t *testing.T) {
	t.Parallel()
	t.Run("Should create new currencies", func(t *testing.T) {
		t.Parallel()
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		require.NoError(t, err)
		gormDB, err := gorm.Open(postgres.New(postgres.Config{
			Conn: db,
		}), &gorm.Config{})
		require.NoError(t, err)

		s := NewRepository(gormDB)
		c := make(internal.Currencies, 0)
		currencyToCreate := internal.Currency{
			Code:  faker.Word(),
			Value: 10,
		}
		c = append(c, currencyToCreate)
		m := internal.MetaData{LastUpdateAt: time.Now()}
		assert.NoError(t, err)
		q := `INSERT INTO "sql_currency" ("created_at","updated_at","deleted_at","code","value","last_updated_at") VALUES ($1,$2,$3,$4,$5,$6)` // nolint:lll
		mock.ExpectBegin()
		mock.ExpectExec(q).
			WithArgs(AnyTime{}, AnyTime{}, nil, currencyToCreate.Code, currencyToCreate.Value, m.LastUpdateAt).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()
		err = s.CreateCurrencies(c, m)
		assert.NoError(t, mock.ExpectationsWereMet())
		assert.NoError(t, err)
	})
}

func TestRepository_FindWithDate(t *testing.T) {
	t.Parallel()
	t.Run("Should get elements of currency by date", func(t *testing.T) {
		t.Parallel()
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		require.NoError(t, err)
		gormDB, err := gorm.Open(postgres.New(postgres.Config{
			Conn: db,
		}), &gorm.Config{})
		code, err := internal.NewCode(faker.Word())
		assert.NoError(t, err)
		t1, errF2 := internal.NewTimeFilter("2022-11-28T20:15:00")
		assert.NoError(t, errF2)
		t2, errF1 := internal.NewTimeFilter("2022-12-28T20:15:00")
		assert.NoError(t, errF1)
		require.NoError(t, err)
		q := `SELECT * FROM "sql_currency" WHERE (code = $1 AND last_updated_at >= $2 AND last_updated_at < $3)
AND "sql_currency"."deleted_at" IS NULL`
		mock.ExpectQuery(q).
			WithArgs(string(code), time.Time(t1), time.Time(t2)).
			WillReturnRows(sqlmock.NewRows([]string{"code", "value", "last_updated_at"}).
				AddRow(string(code), 0, time.Now().UTC()))
		s := NewRepository(gormDB)

		_, err = s.FindWithDate(code, t1, t2)
		assert.NoError(t, mock.ExpectationsWereMet())
		assert.NoError(t, err)
	})

	t.Run("Should get elements of currency", func(t *testing.T) {
		t.Parallel()
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		require.NoError(t, err)
		gormDB, err := gorm.Open(postgres.New(postgres.Config{
			Conn: db,
		}), &gorm.Config{})
		code, err := internal.NewCode(faker.Word())
		assert.NoError(t, err)
		t1, errF2 := internal.NewTimeFilter("2022-11-28T20:15:00")
		assert.NoError(t, errF2)
		t2, errF1 := internal.NewTimeFilter("2022-12-28T20:15:00")
		assert.NoError(t, errF1)
		require.NoError(t, err)
		_, _ = t1, t2
		q := `SELECT * FROM "sql_currency" WHERE "sql_currency"."deleted_at" IS NULL`
		mock.ExpectQuery(q).
			WillReturnRows(sqlmock.NewRows([]string{"code", "value", "last_updated_at"}).
				AddRow(string(code), 0, time.Now().UTC()))
		s := NewRepository(gormDB)

		_, err = s.Find()
		assert.NoError(t, mock.ExpectationsWereMet())
		assert.NoError(t, err)
	})
}

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)

	return ok
}
