package usecase

import (
	"distancing-detect-backend/internal/entity"
	"time"
)

func (u *Usecase) NewViolation(class string, totalViolations int, imageLink string) error {
	timestamp := time.Now()
	violation := entity.NewViolation(class, timestamp, totalViolations, imageLink)
	return u.repository.Create(violation)
}

func (u *Usecase) GetViolationsOfClass(class string) ([]*entity.ViolationData, error) {
	return u.repository.GetByClass(class)
}

func (u *Usecase) GetViolations() ([]*entity.ViolationData, error) {
	return u.repository.List()
}
