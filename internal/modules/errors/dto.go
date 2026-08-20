package errors

import (
	"time"

	"github.com/ThatSoftwareCompany/template-go-api/internal/platform/errstore"
)

type ListRequest struct {
	Endpoint string
	Limit    int
}

type ListResponse struct {
	Items []errstore.ErrorEvent `json:"items"`
}

type ErrorEvent struct {
	OccurredAt    time.Time `json:"occurred_at"`
	CorrelationID string    `json:"correlation_id"`
	Method        string    `json:"method"`
	Path          string    `json:"path"`
	Endpoint      string    `json:"endpoint"`
	StatusCode    int       `json:"status_code"`
	ErrorCode     string    `json:"error_code"`
	Message       string    `json:"message"`
}
