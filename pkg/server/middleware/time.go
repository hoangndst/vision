package middleware

import (
	"context"
	"net/http"
	"time"
)

var StartTimeKey = &contextKey{"startTime"}

func Time(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		if GetStartTime(ctx).IsZero() {
			ctx = context.WithValue(ctx, StartTimeKey, time.Now())
		}
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetStartTime(ctx context.Context) time.Time {
	if ctx == nil {
		return time.Time{}
	}
	if startTime, ok := ctx.Value(StartTimeKey).(time.Time); ok {
		return startTime
	}
	return time.Time{}
}
