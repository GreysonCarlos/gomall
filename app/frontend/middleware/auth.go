package middleware

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

type SessionUerIdKey string

const SessionUerId SessionUerIdKey = "user_id"

func GlobalAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		session := sessions.Default(c)
		ctx = context.WithValue(ctx, SessionUerId, session.Get("user_id"))
		c.Next(ctx)
	}
}