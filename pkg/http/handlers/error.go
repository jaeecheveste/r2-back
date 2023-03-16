package handler

import "errors"

var (
	errQueryParamInvalid  = errors.New("query param invalid")
	errQueryParamNotFound = errors.New("query param not found")
)
