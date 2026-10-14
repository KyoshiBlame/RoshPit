package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/KyoshiBlame/RoshPit/event-api/internal/core/logger"
	"github.com/KyoshiBlame/RoshPit/event-api/internal/core/repository"
	"github.com/KyoshiBlame/RoshPit/event-api/internal/core/transport/http/server"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var timeZone = time.UTC

func main() {

	time.Local = timeZone

	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	defer cancel()

	logger, err := logger.NewLogger(logger.NewMustConfig())
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to init logger:%v\n", err)
		os.Exit(1)
	}

	defer logger.Close()

	logger.Debug("application time zone", zap.Any("zone", timeZone))

	logger.Debug("initiliazling postgres connection pool")
	pool, err := repository.NewPool(ctx, repository.NewConfigMust())
	if err != nil {
		logger.Error("failed to init postgres pool", zap.Error(err))
		os.Exit(1)
	}

	defer pool.Close()

	HTTPServer, err := server.NewHTTPServer(
		server.NewConfigMust(),
		logger,
		gin.Recovery(),
	)

	if err != nil {
		logger.Error("failed to init HTTP server", zap.Error(err))
		os.Exit(1)
	}

	if err := HTTPServer.Start(ctx); err != nil {
		logger.Error("failed to start HTTP server", zap.Error(err))
		os.Exit(1)
	}

}
