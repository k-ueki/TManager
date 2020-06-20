package controller

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/k-ueki/tmanager/server/domain/model"
	"github.com/k-ueki/tmanager/server/usecase"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type (
	FollowerCommandController interface {
		Init(ctx echo.Context) error
	}

	followerCommandController struct {
		usecase.FollowerCommandService
	}
)

func NewFollowerCommandController(db *gorm.DB) FollowerCommandController {
	return &followerCommandController{usecase.NewFollowerCommandUseCase(db)}
}

func (c *followerCommandController) Init(ctx echo.Context) error {
	user := &model.User{
		ID:   1,
		Name: "Hoge",
	}
	if err := c.FollowerCommandService.Init(user); err != nil {
		return errors.Wrap(err, "failed to init")
	}
	return ctx.JSON(http.StatusOK, "success")
}
