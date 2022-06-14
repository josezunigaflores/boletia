package command

import (
	"boletia/internal/utils"
	"context"
)

type Bus interface {
	// Dispatch is the method used to dispatch new commands.
	Dispatch(context.Context, Command) (utils.Response, error)
	// Register is the method used to register a new command handler.
	Register(Type, Handler)
}

//go:generate mockery --case=snake --outpkg=commandmocks --output=commandmocks -name=bus

// Type represents an application command type.
type Type string

type Command interface {
	Type() Type
}

type Handler interface {
	Handle(context.Context, Command) (utils.Response, error)
}
