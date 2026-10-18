package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	core_logger "github.com/KyoshiBlame/RoshPit/event-api/internal/core/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type HTTPServer struct {
	config Config
	log    *core_logger.Logger
	router *gin.Engine
}

func NewHTTPServer(
	config Config,
	log *core_logger.Logger,
	middlewares ...gin.HandlerFunc,
) (*HTTPServer, error) {

	router := gin.New()

	router.Use(middlewares...)

	return &HTTPServer{
		config: config,
		log:    log,
		router: router,
	}, nil
}

func (s *HTTPServer) Start(
	ctx context.Context,
) error {

	server := &http.Server{
		Addr:    s.config.Addr,
		Handler: s.router,
	}

	ch := make(chan error, 1)

	go func() {
		defer close(ch)

		s.log.Warn("Start HTTP server", zap.String("addr", s.config.Addr))

		err := server.ListenAndServe()

		if !errors.Is(err, http.ErrServerClosed) {
			ch <- err
		}

	}()

	select {
	case err := <-ch:
		if err != nil {
			return fmt.Errorf("Listen and Serve HTTP error:%w", err)
		}
	case <-ctx.Done():
		s.log.Warn("ShutDown HTTP server...")

		shutdownCtx, cancel := context.WithTimeout(
			context.Background(),
			s.config.ShutDown,
		)

		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			_ = server.Close()

			return fmt.Errorf("shutDown HTTP server: %w", err)
		}

		s.log.Warn("HTTP server stopped")
	}

	return nil
}

func (s *HTTPServer) Router() *gin.Engine {
	return s.router
}
