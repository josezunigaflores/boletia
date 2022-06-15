package calls

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

var (
	ErrCreateCall = errors.New("error creating the call")
)

type Repository struct {
	db *gorm.DB
}

func NewCallRepository(db *gorm.DB) Repository {
	return Repository{db: db}
}

func (r *Repository) CreateCall(dateExecuteTime time.Time, duration time.Duration, status string) error {
	tx := r.db.Begin()
	defer tx.Rollback()
	if err := tx.Table(sqlTableName).Create(&SQLCall{
		DateExecuteTime: dateExecuteTime,
		Duration:        duration,
		Status:          status,
	}).Error; err != nil {
		return fmt.Errorf("%w:%s", ErrCreateCall, err)
	}
	return tx.Commit().Error
}
