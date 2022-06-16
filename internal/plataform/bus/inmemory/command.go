package inmemory

import (
	"boletia/internal/utils"
	"context"

	"boletia/kit/command"
)

type CommandBus struct {
	handlers map[command.Type]command.Handler
}

func NewCommandBus() *CommandBus {
	return &CommandBus{
		handlers: make(map[command.Type]command.Handler),
	}
}

func (b *CommandBus) Dispatch(ctx context.Context, cmd command.Command) (utils.Response, error) {
	handler, ok := b.handlers[cmd.Type()]
	if !ok {
		return nil, nil
	}

	return handler.Handle(ctx, cmd)
}

func (b *CommandBus) Register(cmdType command.Type, handler command.Handler) {
	b.handlers[cmdType] = handler
}
