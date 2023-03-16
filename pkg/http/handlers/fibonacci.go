package handler

import (
	"net/http"

	"github.com/jaeecheveste/r2-back/pkg/domain"

	"github.com/labstack/echo/v4"
)

type FibonacciHandler struct {
	fibService domain.FibonacciService
}

func NewFibonacciHandler(fibService domain.FibonacciService) *FibonacciHandler {
	return &FibonacciHandler{fibService}
}

func (fh *FibonacciHandler) GetSpiral(c echo.Context) error {
	rows, err := getInt32QueryParam(c, "rows")
	if err != nil {
		return err
	}

	cols, err := getInt32QueryParam(c, "cols")
	if err != nil {
		return err
	}

	spiral, err := fh.fibService.GetSpiral(*rows, *cols)
	if err != nil {
		return handleServiceError(c, err)
	}

	return c.JSON(http.StatusOK, spiral)
}
