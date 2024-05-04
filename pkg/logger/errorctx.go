package logger

import (
	"context"
	"errors"
)

type errorWithCtx struct {
	next   error
	logCtx map[string]any
}

func (e *errorWithCtx) Error() string {
	return e.next.Error()
}

func ErrorCtx(ctx context.Context, err error) context.Context {
	var e *errorWithCtx
	if errors.As(err, &e) {
		return context.WithValue(ctx, LogKey, e.logCtx)
	}

	return ctx
}

func WrapError(ctx context.Context, err error) error {
	ctx = ErrorCtx(ctx, err)

	c := make(map[string]any)
	if x, ok := ctx.Value(LogKey).(map[string]any); ok {
		c = x
	}

	return &errorWithCtx{
		next:   err,
		logCtx: c,
	}
}
