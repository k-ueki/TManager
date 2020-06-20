package presistence

import (
	"github.com/jinzhu/gorm"
	"github.com/k-ueki/tmanager/server/domain/repository"
	gormbulk "github.com/t-tiger/gorm-bulk-insert"
)

type (
	followerCommandPersistence struct {
		DB *gorm.DB
	}
)

func NewFollowerCommandPersistence(db *gorm.DB) repository.FollowerCommandRepository {
	return &followerCommandPersistence{db}
}

func (p *followerCommandPersistence) Reset(inp []interface{}) error {
	if err := p.DB.Exec("TRUNCATE user_follower_tid").Error; err != nil {
		return err
	}

	if err := gormbulk.BulkInsert(p.DB.Table("user_follower_tid"), inp, 3000); err != nil {
		return err
	}

	return nil
}
