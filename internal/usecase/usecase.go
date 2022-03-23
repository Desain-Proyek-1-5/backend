package usecase

import "distancing-detect-backend/internal/repository"

type Usecase struct {
	repository *repository.Repository
}

func NewService(repository repository.Repository) *Usecase {
	return &Usecase{
		repository: &repository,
	}

}
