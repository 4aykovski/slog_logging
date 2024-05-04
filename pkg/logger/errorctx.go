package logger

import (
	"context"
	"errors"
)

type errorWithCtx struct {
	next   error
	logCtx LogCtx
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

	newCtx := LogCtx{Args: make(map[string]any)}
	if prevCtx, ok := ctx.Value(LogKey).(LogCtx); ok {
		for k, v := range prevCtx.Args {
			newCtx.Args[k] = v
		}
	}

	return &errorWithCtx{
		next:   err,
		logCtx: newCtx,
	}
}
