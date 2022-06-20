package currency

import (
	"boletia/internal"
	"boletia/internal/currency"
	"boletia/internal/utils"
	"boletia/kit/command"

	"github.com/gin-gonic/gin"
)

// GetCurrencies is a controller dedicate to getting the currencies.
// @Summary Getting all currency or specific currency.
// @Description the source find within all currencies and returns these.
// @Tags         currencies
// @Accept       json
// @Produce      json
// @Success  201  {object} ResponseCurrencies true "Response general"
// @Failure  400 {object} utils.HTTPResponse true "Response with error field"
// @Failure  500 {object} utils.HTTPResponse true  "Response for any error in server"
// @Router       /api/v1/currencies/:id [GET].
func GetCurrencies(bus command.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c := ctx.Param("id")
		find := ctx.Query("finit")
		fend := ctx.Query("fend")
		// is not necessary check the error because.
		// the struct that returns this function content the error.
		response, _ := bus.Dispatch(ctx, currency.NewCurrencyCommand(c, find, fend)) // nolint:errcheck
		data, ok := response.Data().(internal.Currencies)
		if !ok {
			data = nil
		}

		ctx.JSON(500, ResponseCurrencies{
			HTTPResponse: utils.HTTPResponse{
				Code:    response.Code(),
				Message: response.Message(""),
				Error:   response.Error(),
			},
			Data: data,
		})
	}
}

type ResponseCurrencies struct {
	utils.HTTPResponse
	Data internal.Currencies
}
