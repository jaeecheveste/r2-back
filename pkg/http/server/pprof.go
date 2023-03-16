package server

import (
	"net/http"
	"net/http/pprof"

	"github.com/labstack/echo/v4"
)

// pprofPrefix path definition for pprof execution.
const (
	pprofPrefix = "r2/v1/debug/pprof"
)

// handlerPprof returns the prefix astraction.
func getPrefix(prefixs ...string) string {
	if len(prefixs) > 0 {
		return prefixs[0]
	}
	return pprofPrefix
}

// handlerPprof requests for each handler function.
func handlerPprof(h http.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		h.ServeHTTP(c.Response().Writer, c.Request())
		return nil
	}
}

// RegisterPprof register a middleware for net/http/pprof.
func RegisterPprof(e *echo.Echo, prefixOptions ...string) {
	pprofPrefix := getPrefix(prefixOptions...)
	pprofRouter := e.Group(pprofPrefix)
	{
		pprofRouter.GET("/", handlerPprof(pprof.Index))
		pprofRouter.GET("/allocs", handlerPprof(pprof.Handler("allocs").ServeHTTP))
		pprofRouter.GET("/block", handlerPprof(pprof.Handler("block").ServeHTTP))
		pprofRouter.GET("/cmdline", handlerPprof(pprof.Cmdline))
		pprofRouter.GET("/goroutine", handlerPprof(pprof.Handler("goroutine").ServeHTTP))
		pprofRouter.GET("/heap", handlerPprof(pprof.Handler("heap").ServeHTTP))
		pprofRouter.GET("/mutex", handlerPprof(pprof.Handler("mutex").ServeHTTP))
		pprofRouter.GET("/profile", handlerPprof(pprof.Profile))
		pprofRouter.POST("/symbol", handlerPprof(pprof.Symbol))
		pprofRouter.GET("/symbol", handlerPprof(pprof.Symbol))
		pprofRouter.GET("/threadcreate", handlerPprof(pprof.Handler("threadcreate").ServeHTTP))
		pprofRouter.GET("/trace", handlerPprof(pprof.Trace))
	}
}
