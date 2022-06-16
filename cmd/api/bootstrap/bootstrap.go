package bootstrap

import (
	"boletia/internal"
	"boletia/internal/config"
	"boletia/internal/currency"
	"boletia/internal/plataform/bus/inmemory"
	"boletia/internal/plataform/http"
	"boletia/internal/plataform/server"
	"boletia/internal/plataform/storage/postgres/calls"
	currencyRepository "boletia/internal/plataform/storage/postgres/currency"
	"boletia/internal/plataform/storage/postgres/migrator"
	"boletia/internal/schedule"
	"context"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

func Run() error {
	if err := config.Init(); err != nil {
		return err
	}
	var (
		commandBus = inmemory.NewCommandBus()
		eventBus   = inmemory.NewEventBus()
	)
	cnf := config.Config

	postgresURI := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", config.Config.DBHost, config.Config.DBPort, config.Config.DBUser, config.Config.DBPass, config.Config.DBName)
	db, err := gorm.Open(postgres.Open(postgresURI), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: true,
			NameReplacer:  nil,
			NoLowerCase:   false,
		},
	})
	if err != nil {
		return err
	}

	initMigrator := migrator.NewInitialsRepository(db)
	if err := initMigrator.CreateTables(); err != nil {
		return err
	}

	currencyHTPP := http.NewRepositoryCurrency(cnf.PathCurrency, cnf.TimeOut)
	currencyRep := currencyRepository.NewRepository(db)
	serviceCurrency := currency.NewServiceCurrency(&currencyRep)
	callRepository := calls.NewCallRepository(db)
	cmd := currency.NewCurrencyHandler(serviceCurrency)
	timer := schedule.NewServiceSchedule(&currencyHTPP, &currencyRep, cnf.Schedule, cnf.TimeOut, eventBus)
	go timer.Do()
	eventBus.Subscribe(internal.CurrencyFailEventType, schedule.NewEvent(&callRepository))
	commandBus.Register(currency.CurrencyCommandType, cmd)
	host, port, shutdownTimeout := cnf.Host, cnf.Port, time.Duration(cnf.ShutdownTimeout)*time.Second
	ctx, srv := server.New(context.Background(), host, port, shutdownTimeout, commandBus)

	return srv.Run(ctx)
}
