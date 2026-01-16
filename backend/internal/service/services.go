package service

import (
	"github.com/Sameer16536/Go_BoilerPlate_Code/backend/internal/lib/job"
	"github.com/Sameer16536/Go_BoilerPlate_Code/backend/internal/repository"
	"github.com/Sameer16536/Go_BoilerPlate_Code/backend/internal/server"
)

type Services struct {
	Auth *AuthService
	Job  *job.JobService
}

func NewServices(s *server.Server, repos *repository.Repositories) (*Services, error) {
	authService := NewAuthService(s)

	return &Services{
		Job:  s.Job,
		Auth: authService,
	}, nil
}
