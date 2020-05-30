package repository

import (
	"github.com/k-ueki/tmanager/server/domain/entity"
)

type (
	FollowerRepository interface {
		FindByUserID(userID int) (*entity.User, error)
		FindFollowerByUserID(userID int) ([]*entity.User, error)
		FindFollowByUserID(userID int) ([]*entity.User, error)
		FindUnrequitedUser(userID int) ([]*entity.User, error)
	}
)
