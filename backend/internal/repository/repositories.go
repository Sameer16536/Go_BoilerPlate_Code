package repository

import "github.com/Sameer16536/Go_BoilerPlate_Code/backend/internal/server"

type Repositories struct{}

func NewRepositories(s *server.Server) *Repositories {
	return &Repositories{}
}
