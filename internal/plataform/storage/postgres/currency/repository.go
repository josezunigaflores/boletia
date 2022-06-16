package currency

import (
	"boletia/internal"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

var (
	ErrCreatingCurrencies = errors.New("error creating the currencies")
)

type Repository struct {
	db *gorm.DB
	tx *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{db: db}
}

func (r *Repository) CreateCurrencies(currencies internal.Currencies, data internal.MetaData) error {
	r.tx = r.db.Begin()
	defer r.tx.Rollback()
	for _, c := range currencies {
		if err := r.CreateCurrency(c.Code, c.Value, data.LastUpdateAt); err != nil {
			return fmt.Errorf("%w:%s", ErrCreatingCurrencies, err)
		}
	}

	return r.tx.Commit().Error
}

func (r *Repository) CreateCurrency(code string, value float64, lastUpdatedAt time.Time) error {
	return r.tx.Table(sqlTableName).Create(&SQLCurrency{
		Code:          code,
		Value:         value,
		LastUpdatedAt: lastUpdatedAt,
	}).Error
}

func (r *Repository) FindWithDate(code internal.Code, finit, fend internal.FilterTime) (internal.Currencies, error) {
	currencies := make(Currencies, 0)
	err := r.db.Table(sqlTableName).
		Where("code = ? AND  last_updated_at >= ? AND last_updated_at < ?", code, finit, fend).Find(&currencies).Error
	if err != nil {
		return nil, err
	}

	return currencies.ToCurrencies()
}

func (r *Repository) Find() (internal.Currencies, error) {
	currencies := make(Currencies, 0)
	err := r.db.Table(sqlTableName).Find(&currencies).Error
	if err != nil {
		return nil, err
	}
	return currencies.ToCurrencies()
}
