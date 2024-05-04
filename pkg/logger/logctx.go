package logger

import "context"

func AddArg(ctx context.Context, key string, value any) context.Context {
	return AddArgs(ctx, map[string]any{key: value})
}

func AddArgs(ctx context.Context, args map[string]any) context.Context {
	newArgs := make(map[string]any)
	if prevArgs, ok := ctx.Value(LogKey).(map[string]any); ok {
		newArgs = prevArgs
	}

	for k, v := range args {
		newArgs[k] = v
	}

	return context.WithValue(ctx, LogKey, newArgs)
}

func AddComponent(ctx context.Context, component string) context.Context {
	return AddArg(ctx, "component", component)
}

func AddUserId(ctx context.Context, id string) context.Context {
	return AddArg(ctx, "user_id", id)
}
