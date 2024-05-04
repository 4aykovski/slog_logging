package service

import (
	"context"
	"fmt"

	"github.com/4aykovski/slog_logging/internal/repository"
	"github.com/4aykovski/slog_logging/pkg/logger"
)

func Test(ctx context.Context) (string, error) {
	ctx = logger.AddUserId(ctx, "1")

	err := repository.Test(ctx)
	if err != nil {
		return "", logger.WrapError(ctx, fmt.Errorf("service error: %w", err))
	}
	return "", nil
}
