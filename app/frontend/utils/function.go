package utils

import "context"

func GetUserIdFromCtx(ctx context.Context) uint32 {
	userId := ctx.Value(SessionUerId)
	if userId == nil {
		return 0
	}
	return userId.(uint32)
}