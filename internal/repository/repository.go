package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/4aykovski/slog_logging/pkg/logger"
)

func Test(ctx context.Context) error {
	logArgs := make(map[string]any)
	logArgs["key"] = "value"

	ctx2 := logger.AddArgs(ctx, logArgs)
	ctx2 = logger.AddArg(ctx2, "key2", "value2")

	err := fmt.Errorf("repository error: %w", errors.New("123"))
	if err != nil {
		return logger.WrapError(ctx2, err)
	}

	return nil
}
