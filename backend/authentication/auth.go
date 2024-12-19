package auth

import (
	"context"
	"errors"
)

type contextKey string

const userIDKey = contextKey("userID")

func AddUserIDToContext(ctx context.Context, userID int) context.Context {
	return context.WithValue(ctx, userIDKey, userID)
}

func GetUserIDFromContext(ctx context.Context) (int, error) {
	userID, ok := ctx.Value(userIDKey).(int)
	if !ok {
		return 0, errors.New("userID not found in context")
	}
	return userID, nil
}
