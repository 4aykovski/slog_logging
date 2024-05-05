package main

import (
	"context"

	"github.com/4aykovski/slog_logging/internal/service"
	"github.com/4aykovski/slog_logging/pkg/logger"
)

func main() {
	logger.InitLogging("local")

	ctx := context.Background()
	logArgs := make(map[string]any)
	logArgs["component"] = "main"
	ctx = logger.AddArgs(ctx, logArgs)

	_, err := service.Test(ctx)
	if err != nil {
		logger.ErrCtx(logger.ErrorCtx(ctx, err), err.Error())
	}

	logger.InfoCtx(ctx, "test")
}
