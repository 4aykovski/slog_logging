package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/4aykovski/slog_logging/pkg/logger"
)

func Test(ctx context.Context) error {
	logArgs := make(map[string]any)
	logArgs["component"] = "repo"
	ctx = logger.AddArgs(ctx, logArgs)

	logger.InfoCtx(ctx, "test")

	ctx = logger.AddArg(ctx, "key2", "value2")
	ctx = logger.AddComponent(ctx, "repo")

	err := fmt.Errorf("repository error: %w", errors.New("123"))
	if err != nil {
		return logger.WrapError(ctx, err)
	}

	return nil
}
