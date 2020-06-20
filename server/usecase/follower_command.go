package usecase

import (
	"github.com/jinzhu/gorm"
	"github.com/k-ueki/tmanager/server/domain/entity"
	"github.com/k-ueki/tmanager/server/domain/model"
	"github.com/k-ueki/tmanager/server/domain/repository"
	"github.com/k-ueki/tmanager/server/infra/presistence"
)

type (
	FollowerCommandService interface {
		Init(user *model.User) error
	}

	followerCommandUseCase struct {
		FollowerQueryUseCase
		repository.FollowerCommandRepository
	}
)

func NewFollowerCommandUseCase(db *gorm.DB) FollowerCommandService {
	return &followerCommandUseCase{
		NewFollowerQueryUseCase(db),
		presistence.NewFollowerCommandPersistence(db),
	}
}

func (s *followerCommandUseCase) Init(user *model.User) error {
	follower, err := s.FollowerQueryUseCase.GetFollowersIDsFromTwitterAPI()
	if err != nil {
		return err
	}

	var list []interface{}
	for _, id := range follower.IDs {
		list = append(list, entity.UserFollowerTid{
			UserID:      uint(user.ID),
			FollowerTid: uint64(id),
		})
	}

	if err := s.FollowerCommandRepository.Reset(list); err != nil {
		return err
	}

	return nil
}
