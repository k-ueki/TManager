package presistence

import (
	"github.com/jinzhu/gorm"
	"github.com/k-ueki/tmanager/server/domain/repository"
)

type (
	followerCommandPersistence struct {
		DB *gorm.DB
	}
)

func NewFollowerCommandPersistence(db *gorm.DB) repository.FollowerCommandRepository {
	return &followerCommandPersistence{db}
}
