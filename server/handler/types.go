package handler

import (
	"fmt"
	"net/http"
	"time"
)

type Response struct {
	Success   bool       `json:"success" yaml:"success"`
	Message   string     `json:"message" yaml:"message"`
	Data      any        `json:"data,omitempty" yaml:"data,omitempty"`
	TraceID   string     `json:"trace_id,omitempty" yaml:"trace_id,omitempty"`
	StartTime *time.Time `json:"start_time,omitempty" yaml:"start_time,omitempty"`
	EndTime   *time.Time `json:"end_time,omitempty" yaml:"end_time,omitempty"`
	CostTime  Duration   `json:"cost_time,omitempty" yaml:"cost_time,omitempty"` // Time taken to process the request.
}

type Duration time.Duration

func (resp *Response) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (d Duration) MarshalJSON() (b []byte, err error) {
	// Format the duration as a string.
	return []byte(fmt.Sprintf(`"%s"`, time.Duration(d).String())), nil
}
