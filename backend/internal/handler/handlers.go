package handler

import (
	"github.com/Sameer16536/Go_BoilerPlate_Code/backend/internal/server"
	"github.com/Sameer16536/Go_BoilerPlate_Code/backend/internal/service"
)

type Handlers struct {
	Health  *HealthHandler
	OpenAPI *OpenAPIHandler
}

func NewHandlers(s *server.Server, services *service.Services) *Handlers {
	return &Handlers{
		Health:  NewHealthHandler(s),
		OpenAPI: NewOpenAPIHandler(s),
	}
}
