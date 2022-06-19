package currency

import (
	"boletia/internal"
	"boletia/internal/utils"
	"boletia/kit/command"
	"context"
	"errors"
)

const CurrencyCommandType command.Type = "command.currency.creating"

type Command struct {
	code        string
	finit, fend string
}

func (c Command) Type() command.Type {
	return CurrencyCommandType
}

func NewCurrencyCommand(code string, finit string, fend string) *Command {
	return &Command{code: code, finit: finit, fend: fend}
}

type CurrencyHandler struct {
	service Service
}

func NewCurrencyHandler(service Service) *CurrencyHandler {
	return &CurrencyHandler{service: service}
}

var errUnexpected = errors.New("unexpected command")

func (ch CurrencyHandler) Handle(ctx context.Context, cmd command.Command) (response utils.Response, err error) {
	cc, ok := cmd.(*Command)
	if !ok {
		err := errUnexpected
		response = utils.NewInternalErrResponse(err, err)

		return response, err
	}
	currencies, err := ch.service.FindCurrency(cc.code, cc.finit, cc.fend)
	// defines bad request
	if errors.Is(err, internal.ErrBadCode) || errors.Is(err, internal.ErrBadTimeFilter) {
		response = utils.NewBadRequest(err)

		return response, err
	}
	if err != nil {
		response = utils.NewInternalErrResponse(err, err)

		return response, err
	}
	responseSuccess := NewSuccessCurrency(currencies)

	return responseSuccess, nil
}
