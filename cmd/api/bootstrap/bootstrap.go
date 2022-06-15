package bootstrap

import (
	"boletia/internal/config"
	"boletia/internal/plataform/server"
	"context"
	"time"
)

func Run() error {
	if err := config.Init(); err != nil {
		return err
	}
	cnf := config.Config
	host, port, shutdownTimeout := cnf.Host, cnf.Port, time.Duration(cnf.ShutdownTimeout)*time.Second
	ctx, srv := server.New(context.Background(), host, port, shutdownTimeout)

	return srv.Run(ctx)
}
