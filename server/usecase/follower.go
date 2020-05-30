package usecase

import (
	"github.com/k-ueki/tmanager/server/config"
	"github.com/pkg/errors"

	"github.com/jinzhu/gorm"
	"github.com/k-ueki/tmanager/server/domain/entity"
	"github.com/k-ueki/tmanager/server/domain/model"
	"github.com/k-ueki/tmanager/server/domain/repository"
	"github.com/k-ueki/tmanager/server/infra/presistence"
)

type (
	FollowerUseCase interface {
		Show(userID int) (*entity.User, error)
		ListFollower(user *model.User) ([]*entity.User, error)
		ListFollow(user *model.User) ([]*entity.User, error)
		ListUnrequitedUsers(user *model.User) ([]*entity.User, error)
		GetFollowersFromTwitterAPI() (*model.UserFromTwitterAPI, error)
		GetFollowersIDsFromTwitterAPI() (*model.UserIDs, error)
	}

	followerUseCase struct {
		config.TwitterAPIClient
		repository.FollowerRepository
	}
)

func NewFollowerUseCase(db *gorm.DB) FollowerUseCase {
	return &followerUseCase{
		*config.NewTwitterAPIClient(),
		presistence.NewFollowerPersistence(db),
	}
}

// The people following the User
func (u *followerUseCase) Show(userID int) (*entity.User, error) {
	user, err := u.FollowerRepository.FindByUserID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// The people following the User
func (u *followerUseCase) ListFollower(user *model.User) ([]*entity.User, error) {
	followers, err := u.FollowerRepository.FindFollowerByUserID(user.ID)
	if err != nil {
		return nil, err
	}
	return followers, nil
}

// The people the User follow
func (u *followerUseCase) ListFollow(user *model.User) ([]*entity.User, error) {
	followers, err := u.FollowerRepository.FindFollowByUserID(user.ID)
	if err != nil {
		return nil, err
	}
	return followers, nil
}

func (u *followerUseCase) ListUnrequitedUsers(user *model.User) ([]*entity.User, error) {
	unrequitedUsers, err := u.FollowerRepository.FindUnrequitedUser(user.ID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to list the people I follow")
	}
	return unrequitedUsers, nil
}

func (u *followerUseCase) GetFollowersFromTwitterAPI() (*model.UserFromTwitterAPI, error) {
	followers := model.UserFromTwitterAPI{}
	if err := u.TwitterAPIClient.HTTPRequest(config.PathToGetFollowers, &followers); err != nil {
		return nil, err
	}
	return &followers, nil
}

func (u *followerUseCase) GetFollowersIDsFromTwitterAPI() (*model.UserIDs, error) {
	ids := model.UserIDs{}
	if err := u.TwitterAPIClient.HTTPRequest(config.PathToGetFollowersIDs, &ids); err != nil {
		return nil, err
	}
	return &ids, nil
}
