package usecase

import (
	"fmt"

	"github.com/k-ueki/tmanager/server/config"
	"github.com/pkg/errors"

	"github.com/jinzhu/gorm"
	"github.com/k-ueki/tmanager/server/domain/entity"
	"github.com/k-ueki/tmanager/server/domain/model"
	"github.com/k-ueki/tmanager/server/domain/repository"
	"github.com/k-ueki/tmanager/server/infra/presistence"
)

type (
	FollowerQueryUseCase interface {
		Show(userID int) (*entity.User, error)
		ListFollower(user *model.User) ([]*entity.User, error)
		ListFollow(user *model.User) ([]*entity.User, error)
		ListNewFollower(user *model.User) ([]*model.UserInfoFromTwitterAPI, error)
		ListUnrequitedUsers(user *model.User) ([]*entity.User, error)
		GetFollowersFromTwitterAPI() (*model.UserFromTwitterAPI, error)
		GetFollowersIDsFromTwitterAPI() (*model.UserIDs, error)
	}

	followerQueryUseCase struct {
		config.TwitterAPIClient
		repository.FollowerQueryRepository
	}
)

func NewFollowerQueryUseCase(db *gorm.DB) FollowerQueryUseCase {
	return &followerQueryUseCase{
		*config.NewTwitterAPIClient(),
		presistence.NewFollowerQueryPersistence(db),
	}
}

// The people following the User
func (u *followerQueryUseCase) Show(userID int) (*entity.User, error) {
	user, err := u.FollowerQueryRepository.FindByUserID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// The people following the User
func (u *followerQueryUseCase) ListFollower(user *model.User) ([]*entity.User, error) {
	followers, err := u.FollowerQueryRepository.FindFollowerByUserID(user.ID)
	if err != nil {
		return nil, err
	}
	return followers, nil
}

// The people the User follow
func (u *followerQueryUseCase) ListFollow(user *model.User) ([]*entity.User, error) {
	followers, err := u.FollowerQueryRepository.FindFollowByUserID(user.ID)
	if err != nil {
		return nil, err
	}
	return followers, nil
}

func (u *followerQueryUseCase) ListUnrequitedUsers(user *model.User) ([]*entity.User, error) {
	unrequitedUsers, err := u.FollowerQueryRepository.FindUnrequitedUser(user.ID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to list the people I follow")
	}
	return unrequitedUsers, nil
}

func (u *followerQueryUseCase) GetFollowersFromTwitterAPI() (*model.UserFromTwitterAPI, error) {
	followers := model.UserFromTwitterAPI{}
	if err := u.TwitterAPIClient.HTTPRequest(config.PathToGetFollowers, &followers); err != nil {
		return nil, err
	}
	return &followers, nil
}

func (u *followerQueryUseCase) GetFollowersIDsFromTwitterAPI() (*model.UserIDs, error) {
	ids := model.UserIDs{}
	if err := u.TwitterAPIClient.HTTPRequest(config.PathToGetFollowersIDs, &ids); err != nil {
		return nil, err
	}
	return &ids, nil
}

func (u *followerQueryUseCase) ListNewFollower(user *model.User) ([]*model.UserInfoFromTwitterAPI, error) {
	nowFollower, err := u.GetFollowersIDsFromTwitterAPI()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get followersIDs from TwitterAPI")
	}

	original, err := u.FollowerQueryRepository.FindFollowerTwitterIDsByUserID(user.ID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get followersIDs from DB")
	}

	//TODO:entityにconverterがあるのは気持ち悪い
	newFollowers := u.getNewFollowerList(nowFollower.ConvertUserIDsToUint64(), entity.ConvertTwitterIDToUint64(original))

	return u.getUserInfoFromTwitterAPIByTid(newFollowers)
}

func (u *followerQueryUseCase) getNewFollowerList(nows, originals []*uint64) []*uint64 {
	var newFollower []*uint64

	if len(originals) == 0 {
		return nows
	}

	for _, now := range nows {
		for i, original := range originals {
			if *now == *original {
				break
			}
			if i == len(originals)-1 {
				newFollower = append(newFollower, now)
			}
		}
	}
	return newFollower
}

func (u *followerQueryUseCase) getUserInfoFromTwitterAPIByTid(ids []*uint64) ([]*model.UserInfoFromTwitterAPI, error) {
	uids := ""
	for i, id := range ids {
		if i == len(ids)-1 {
			uids += fmt.Sprintf("%d", *id)
			break
		}
		uids += fmt.Sprintf("%d&", *id)
	}

	var uinfos []*model.UserInfoFromTwitterAPI

	err := u.TwitterAPIClient.HTTPRequest(fmt.Sprintf(config.PathToGetFollowerInfo+"?user_id=%s", uids), &uinfos)
	if err != nil {
		return nil, err
	}

	return uinfos, nil
}
