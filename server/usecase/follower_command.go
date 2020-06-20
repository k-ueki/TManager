package usecase

import (
	"github.com/jinzhu/gorm"
	"github.com/k-ueki/tmanager/server/domain/repository"
	"github.com/k-ueki/tmanager/server/infra/presistence"
)

type (
	FollowerCommandService interface {
	}

	followerCommandUseCase struct {
		repository.FollowerCommandRepository
	}
)

func NewFollowerCommandUseCase(db *gorm.DB) FollowerCommandService {
	return &followerCommandUseCase{
		presistence.NewFollowerCommandPersistence(db),
	}
}
