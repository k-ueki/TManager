package controller

import (
	"github.com/jinzhu/gorm"
	"github.com/k-ueki/tmanager/server/usecase"
)

type (
	FollowerCommandController interface {
	}

	followerCommandController struct {
		usecase.FollowerCommandService
	}
)

func NewFollowerCommandController(db *gorm.DB) FollowerCommandController {
	return &followerCommandController{usecase.NewFollowerCommandUseCase(db)}
}
