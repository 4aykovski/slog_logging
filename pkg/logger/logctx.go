package logger

import "context"

type LogCtx struct {
	Args map[string]any
}

func AddArgs(ctx context.Context, args map[string]any) context.Context {
	newCtx := LogCtx{Args: make(map[string]any)}

	if prevCtx, ok := ctx.Value(LogKey).(LogCtx); ok {
		for k, v := range prevCtx.Args {
			newCtx.Args[k] = v
		}
	}

	for k, v := range args {
		newCtx.Args[k] = v
	}

	return context.WithValue(ctx, LogKey, newCtx)
}

func AddArg(ctx context.Context, key string, value any) context.Context {
	return AddArgs(ctx, map[string]any{key: value})
}

func AddComponent(ctx context.Context, component string) context.Context {
	return AddArg(ctx, "component", component)
}

func AddUserId(ctx context.Context, id string) context.Context {
	return AddArg(ctx, "user_id", id)
}
