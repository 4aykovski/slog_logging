package logger

import (
	"context"
	"log/slog"
)

type logCtxKey int

const (
	LogKey logCtxKey = iota
)

type HandlerMiddleware struct {
	next slog.Handler
}

func NewHandleMiddleware(next slog.Handler) *HandlerMiddleware {
	return &HandlerMiddleware{next: next}
}

func (h *HandlerMiddleware) Handle(ctx context.Context, rec slog.Record) error {
	if logArgs, ok := ctx.Value(LogKey).(LogCtx); ok {

		for logKey, logValue := range logArgs.Args {
			rec.Add(logKey, logValue)
		}

	}

	return h.next.Handle(ctx, rec)
}

func (h *HandlerMiddleware) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &HandlerMiddleware{next: h.next.WithAttrs(attrs)}
}

func (h *HandlerMiddleware) WithGroup(name string) slog.Handler {
	return &HandlerMiddleware{next: h.next.WithGroup(name)}
}

func (h *HandlerMiddleware) Enabled(ctx context.Context, rec slog.Level) bool {
	return h.next.Enabled(ctx, rec)
}
