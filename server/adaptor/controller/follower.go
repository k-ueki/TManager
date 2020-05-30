package controller

import (
	"errors"
	"net/http"

	"github.com/k-ueki/tmanager/server/adaptor/input"

	"github.com/jinzhu/gorm"
	"github.com/k-ueki/tmanager/server/domain/model"
	"github.com/k-ueki/tmanager/server/usecase"
	"github.com/labstack/echo/v4"
)

type (
	FollowerController interface {
		Show(ctx echo.Context) error
		ListFollow(ctx echo.Context) error
		ListFollower(ctx echo.Context) error
		ListUnrequited(ctx echo.Context) error
	}

	followerController struct {
		usecase.FollowerUseCase
	}
)

func NewFollowerController(db *gorm.DB) FollowerController {
	return &followerController{usecase.NewFollowerUseCase(db)}
}

func (c *followerController) ListFollow(ctx echo.Context) error {
	user := model.User{
		ID:   1,
		Name: "User1",
	}

	followers, err := c.FollowerUseCase.ListFollow(&user)
	if err != nil {
		return errors.New("failed to lis followers")
	}
	return ctx.JSON(http.StatusOK, followers)
}

func (c *followerController) ListFollower(ctx echo.Context) error {
	user := model.User{
		ID:   1,
		Name: "User1",
	}

	followers, err := c.FollowerUseCase.ListFollower(&user)
	if err != nil {
		return errors.New("failed to lis followers")
	}
	return ctx.JSON(http.StatusOK, followers)
}

func (c *followerController) Show(ctx echo.Context) error {
	param := input.ShowFollowerParam{}
	if err := ctx.Bind(&param); err != nil {
		return ctx.JSON(http.StatusBadRequest, "invalid params")
	}

	followers, err := c.FollowerUseCase.Show(param.UserID)
	if err != nil {
		return errors.New("failed to lis followers")
	}
	return ctx.JSON(http.StatusOK, followers)
}

func (c *followerController) ListUnrequited(ctx echo.Context) error {
	user := model.User{
		ID:   1,
		Name: "User1",
	}

	unrequitedUsers, err := c.FollowerUseCase.ListUnrequitedUsers(&user)
	if err != nil {
		return errors.New("failed to lis followers")
	}
	return ctx.JSON(http.StatusOK, unrequitedUsers)
}
