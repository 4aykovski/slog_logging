package main

import (
	"context"

	"github.com/4aykovski/slog_logging/internal/service"
	"github.com/4aykovski/slog_logging/pkg/logger"
)

func main() {
	logger.InitLogging("local")

	ctx := context.Background()

	ctx = logger.AddComponent(ctx, "main")

	_, err := service.Test(ctx)
	if err != nil {
		logger.ErrCtx(logger.ErrorCtx(ctx, err), err.Error())
	}

	logger.InfoCtx(ctx, "test")
}
