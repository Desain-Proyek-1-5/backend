package httphandler

import (
	"distancing-detect-backend/pkg/database"
	"distancing-detect-backend/pkg/logger"
	"distancing-detect-backend/pkg/router"
)

type HTTPHandlers struct {
	Router   *router.RouterInstance
	Logger   *logger.LoggerInstance
	Database *database.DBInstance
}

func NewHTTPHandler(Router *router.RouterInstance, Database *database.DBInstance, Logger *logger.LoggerInstance) *HTTPHandlers {
	return &HTTPHandlers{Router, Logger, Database}
}

func (h *HTTPHandlers) RegisterHandlers() {
	h.Router.RegisterHandler("/{class}", h.GetViolationsOfClass, "GET")
	h.Router.RegisterHandler("/", h.GetAllViolations, "GET")
}
