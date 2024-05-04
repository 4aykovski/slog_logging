package logger

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"runtime"
	"time"
)

func InitLogging(env string) {
	var handler slog.Handler
	switch env {
	case "local":
		handler = slog.Handler(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level:     slog.LevelDebug,
			AddSource: true,
		}))
	}

	handler = New(handler)
	slog.SetDefault(slog.New(handler))
}

func InfoCtx(ctx context.Context, format string, args ...any) {
	if !slog.Default().Enabled(ctx, slog.LevelInfo) {
		return
	}

	var pcs = getPcs()
	r := slog.NewRecord(time.Now(), slog.LevelInfo, fmt.Sprintf(format, args...), pcs[0])
	_ = slog.Default().Handler().Handle(ctx, r)
}

func ErrCtx(ctx context.Context, format string, args ...any) {
	if !slog.Default().Enabled(ctx, slog.LevelError) {
		return
	}

	var pcs = getPcs()
	r := slog.NewRecord(time.Now(), slog.LevelError, fmt.Sprintf(format, args...), pcs[0])
	_ = slog.Default().Handler().Handle(ctx, r)
}

func WarnCtx(ctx context.Context, format string, args ...any) {
	if !slog.Default().Enabled(ctx, slog.LevelWarn) {
		return
	}

	var pcs = getPcs()
	r := slog.NewRecord(time.Now(), slog.LevelWarn, fmt.Sprintf(format, args...), pcs[0])
	_ = slog.Default().Handler().Handle(ctx, r)
}

func DebugCtx(ctx context.Context, format string, args ...any) {
	if !slog.Default().Enabled(ctx, slog.LevelDebug) {
		return
	}

	var pcs = getPcs()
	r := slog.NewRecord(time.Now(), slog.LevelDebug, fmt.Sprintf(format, args...), pcs[0])
	_ = slog.Default().Handler().Handle(ctx, r)
}

func getPcs() []uintptr {
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:])
	return pcs[:]
}
