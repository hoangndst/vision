package handler

import (
	"context"
	"github.com/go-chi/render"
	"github.com/hoangndst/vision/server/middleware"
	"time"
)

const SuccessMessage = "OK"

func GenerateResponse(ctx context.Context, data any, err error) render.Renderer {
	response := &Response{}
	if err == nil {
		response.Success = true
		response.Message = SuccessMessage
		response.Data = data
	} else {
		response.Success = false
		response.Message = err.Error()
	}
	if traceID := middleware.GetTraceID(ctx); len(traceID) > 0 {
		response.TraceID = traceID
	}

	if startTime := middleware.GetStartTime(ctx); !startTime.IsZero() {
		endTime := time.Now()
		response.StartTime = &startTime
		response.EndTime = &endTime
		response.CostTime = Duration(endTime.Sub(startTime))
	}
	return response
}

// FailureResponse creates a response renderer for a failed request.
func FailureResponse(ctx context.Context, err error) render.Renderer {
	return GenerateResponse(ctx, nil, err)
}

// SuccessResponse creates a response renderer for a successful request.
func SuccessResponse(ctx context.Context, data any) render.Renderer {
	return GenerateResponse(ctx, data, nil)
}
