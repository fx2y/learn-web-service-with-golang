package handlers

import (
	"context"
	"github.com/fx2y/learn-web-service-with-golang/internal/platform/web"
	"github.com/julienschmidt/httprouter"
	"go.opencensus.io/trace"
	"net/http"
)

// Check provides support for orchestration health checks.
type Check struct {
	build string
}

// Health validates the service is healthy and ready to accept requests.
func (c *Check) Health(ctx context.Context, w http.ResponseWriter, r *http.Request, params httprouter.Params) error {
	ctx, span := trace.StartSpan(ctx, "handlers.Check.Health")
	defer span.End()

	health := struct {
		Version string `json:"version"`
		Status  string `json:"status"`
	}{
		Version: c.build,
	}

	health.Status = "ok"
	return web.Respond(ctx, w, health, http.StatusOK)
}
