package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/jaeecheveste/r2-back/pkg/config"
	handler "github.com/jaeecheveste/r2-back/pkg/http/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	swagger "github.com/swaggo/echo-swagger"
)

// Server runs an http server.
type Server struct {
	server *echo.Echo
	port   int
}

// NewServer server instantiation for endpoint exposure and middleware configuration.
func NewServer(serverConfig *config.ServerConfig, fh *handler.FibonacciHandler) *Server {
	//pprofEnable := env.GetBoolEnv("PPROF_ENABLED", false)

	e := echo.New()

	// middleware
	registerMiddlewares(e, false)

	// routes prefix
	path := e.Group("/r2")
	apiv1 := path.Group("/v1")

	// swagger docs route
	apiv1.GET("/swagger/*any", swagger.WrapHandler)

	// healthcheck route
	apiv1.GET("/health", health)

	// fibonacci routes
	apiv1.GET("/spiral", fh.GetSpiral)

	return &Server{server: e, port: serverConfig.ServerPort}
}

func registerMiddlewares(server *echo.Echo, pprofEnable bool) {
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())
	server.Use(middleware.CORS())

	if pprofEnable {
		RegisterPprof(server)
	}
}

// Start runs an http server.
func (s *Server) Start() error {
	port := fmt.Sprintf(":%v", s.port)
	if err := s.server.Start(port); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("server.Start: %w", err)
	}

	return nil
}

// Stop stops an http server.
func (s *Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("server.Shutdown: %w", err)
	}

	return nil
}

func health(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}
