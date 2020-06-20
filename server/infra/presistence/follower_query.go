package presistence

import (
	"github.com/jinzhu/gorm"
	"github.com/k-ueki/tmanager/server/domain/entity"
	"github.com/k-ueki/tmanager/server/domain/repository"
	"github.com/pkg/errors"
)

type (
	followerQueryPersistence struct {
		DB *gorm.DB
	}
)

func NewFollowerQueryPersistence(db *gorm.DB) repository.FollowerQueryRepository {
	return &followerQueryPersistence{db}
}

func (p *followerQueryPersistence) FindByUserID(userID int) (*entity.User, error) {
	var row entity.User
	if err := p.DB.LogMode(true).First(&row, userID).Error; err != nil {
		return nil, err
	}
	return &row, nil
}

func (p *followerQueryPersistence) FindFollowerByUserID(userID int) ([]*entity.User, error) {
	var rows []*entity.User
	if err := p.DB.Joins("JOIN user_follow_user ufu ON user.id = ufu.user_id").Where("ufu.follow_user_id = ?", userID).Find(&rows).Error; err != nil {
		return nil, err
	}
	return rows, nil
}

func (p *followerQueryPersistence) FindFollowByUserID(userID int) ([]*entity.User, error) {
	var rows []*entity.User
	if err := p.DB.Joins("JOIN user_follow_user ufu ON user.id = ufu.follow_user_id").Where("ufu.user_id = ?", userID).Find(&rows).Error; err != nil {
		return nil, err
	}
	return rows, nil
}

func (p *followerQueryPersistence) FindUnrequitedUser(userID int) ([]*entity.User, error) {
	var rows []*entity.User
	if err := p.DB.
		Joins(`JOIN (
			SELECT ufu1.follow_user_id AS unrequited_user_id FROM user_follow_user ufu1 LEFT JOIN (
					SELECT * FROM user_follow_user ufu2
					WHERE ufu2.follow_user_id = ?
				)a ON ufu1.follow_user_id = a.user_id
				WHERE ufu1.user_id = ? AND a.user_id IS NULL
			)b ON user.id = b.unrequited_user_id`, userID, userID).Find(&rows).Error; err != nil {
		return nil, err
	}

	return rows, nil
}

func (p *followerQueryPersistence) FindFollowerTwitterIDsByUserID(userID int) ([]*entity.TwitterID, error) {
	var rows []*entity.TwitterID
	if err := p.DB.Table("user_follower_tid").Where("user_id = ?", userID).Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to find follower TwitterIDs By userID")
	}
	return rows, nil
}
