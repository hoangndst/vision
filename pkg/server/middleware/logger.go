package middleware

import (
	"bytes"
	"context"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/go-chi/httplog/v2"
	"k8s.io/klog/v2"
)

var (
	APILoggerKey       = &contextKey{"logger"}
	RunLoggerKey       = &contextKey{"runLogger"}
	RunLoggerBufferKey = &contextKey{"runLoggerBuffer"}
)

func InitLogger(logFilePath string, name string) *httplog.Logger {
	logWriter, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
	if err != nil {
		if os.IsNotExist(err) {
			logFile := filepath.Dir(logFilePath)
			klog.Infof("Log file path %s does not exist, creating it", logFile)
			if err := os.MkdirAll(logFile, 0o755); err != nil {
				klog.Fatalf("Failed to create log file path %s: %v", logFile, err)
			}
		} else {
			klog.Fatalf("Failed to open log file %s: %v", logFilePath, err)
		}
	}
	logger := httplog.NewLogger(name, httplog.Options{
		LogLevel:        slog.LevelInfo,
		Concise:         true,
		TimeFieldFormat: time.RFC3339,
		Writer:          logWriter,
		RequestHeaders:  true,
		Trace: &httplog.TraceOptions{
			HeaderTrace: TraceIDHeader,
		},
	})
	return logger
}

func InitLoggerBuffer(name string) (*httplog.Logger, *bytes.Buffer) {
	var buffer bytes.Buffer
	logger := httplog.NewLogger(name, httplog.Options{
		LogLevel:        slog.LevelInfo,
		Concise:         true,
		TimeFieldFormat: time.RFC3339,
		Writer:          &buffer,
		RequestHeaders:  true,
		Trace: &httplog.TraceOptions{
			HeaderTrace: TraceIDHeader,
		},
	})
	return logger, &buffer
}

func APILoggerMiddleware(logFile string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			// Retrieve the request ID from the context and create a logger with it.
			if requestID := GetTraceID(ctx); len(requestID) > 0 {
				// Set the output file for klog
				logger := InitLogger(logFile, requestID)
				runLogger, logBuffer := InitLoggerBuffer(requestID)
				ctx = context.WithValue(ctx, APILoggerKey, logger)
				ctx = context.WithValue(ctx, RunLoggerKey, runLogger)
				ctx = context.WithValue(ctx, RunLoggerBufferKey, logBuffer)
			}
			// Continue serving the request with the new context.
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func DefaultLoggerMiddleware(logFile string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		logger := InitLogger(logFile, "DefaultLogger")
		return httplog.RequestLogger(logger)(next)
	}
}
