package handler

import (
	"net/http"
	"strconv"

	"github.com/jaeecheveste/r2-back/pkg/domain"

	"github.com/labstack/echo/v4"
)

// getUintQueryParam returns value of the parameter in unit.
func getInt32QueryParam(c echo.Context, name string) (*int32, error) {
	q := c.QueryParam(name)

	if q == "" {
		return nil, newError(c, http.StatusBadRequest, errQueryParamNotFound)
	}

	queryParam, err := strconv.ParseInt(q, 10, 32)
	if err != nil {
		return nil, newError(c, http.StatusBadRequest, errQueryParamInvalid)
	}

	resp := int32(queryParam)
	return &resp, err
}

// newError returns a new controlled error.
func newError(c echo.Context, code int, err error) *echo.HTTPError {
	return echo.NewHTTPError(code, err.Error())
}

// handleServiceError returns an error according to its type.
func handleServiceError(c echo.Context, err error) *echo.HTTPError {
	switch err.Error() {
	case domain.ErrInternalError.Error():
		return newError(c, http.StatusInternalServerError, err)
	default:
		return newError(c, http.StatusInternalServerError, err)
	}
}
