package middleware

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/go-chi/httplog/v2"
	"github.com/hoangndst/vision/domain/constant"
	"net/http"
	"os"
	"strings"
	"sync/atomic"
)

type (
	contextKeyTraceID int
	contextKeyUserID  string
)

const (
	TraceIDKey contextKeyTraceID = 0
	UserIDKey  contextKeyUserID  = "user_id"
)

var (
	TraceIDHeader = "x-vision-trace"
	UserIDHeader  = "x-vision-user"
)

var (
	prefix string
	reqid  uint64
)

func init() {
	hostname, err := os.Hostname()
	if hostname == "" || err != nil {
		hostname = "localhost"
	}
	var buffer [12]byte
	var b64 string
	for len(b64) < 10 {
		_, err := rand.Read(buffer[:])
		if err != nil {
			return
		}
		b64 = base64.StdEncoding.EncodeToString(buffer[:])
		b64 = strings.NewReplacer("+", "", "/", "").Replace(b64)
	}
	prefix = fmt.Sprintf("%s/%s", hostname, b64[0:10])
}

func TraceID(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		traceID := r.Header.Get(TraceIDHeader)
		if traceID == "" {
			id := atomic.AddUint64(&reqid, 1)
			traceID = fmt.Sprintf("%s-%06d", prefix, id)
		}
		ctx = context.WithValue(ctx, TraceIDHeader, traceID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

func GetTraceID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	if traceID, ok := ctx.Value(TraceIDKey).(string); ok {
		return traceID
	}
	return ""
}

func TraceUserID(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		userID := r.Header.Get(UserIDHeader)
		if userID == "" {
			userID = constant.DefaultUser
		}
		ctx = context.WithValue(ctx, UserIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

func GetTraceUserID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	var logger *httplog.Logger
	if apiLogger, ok := ctx.Value(APILoggerKey).(*httplog.Logger); ok {
		logger = apiLogger
	} else {
		logger = httplog.NewLogger("DefaultLogger")
	}
	if userID, ok := ctx.Value(UserIDKey).(string); ok {
		logger.Info("TraceUserID: ", "user_id", userID)
		return userID
	}
	return ""
}
