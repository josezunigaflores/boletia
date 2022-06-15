package migrator

import (
	"boletia/internal/plataform/storage/postgres/calls"
	"boletia/internal/plataform/storage/postgres/currency"
	"gorm.io/gorm"
)

type InitialsRepository struct {
	db *gorm.DB
}

func NewInitialsRepository(db *gorm.DB) *InitialsRepository {
	return &InitialsRepository{db: db}
}

func (ir *InitialsRepository) CreateTables() error {
	ir.db = ir.db.Begin()
	defer ir.db.Rollback()
	migrator := ir.db.Migrator()
	if !migrator.HasTable(&calls.SQLCall{}) {
		if err := migrator.CreateTable(&calls.SQLCall{}); err != nil {
			return err
		}
	}

	if !migrator.HasTable(&currency.SQLCurrency{}) {
		if err := migrator.CreateTable(&currency.SQLCurrency{}); err != nil {
			return err
		}
	}

	return ir.db.Commit().Error
}
