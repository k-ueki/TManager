package controller

import (
	"net/http"

	"github.com/k-ueki/tmanager/server/adaptor/input"

	"github.com/jinzhu/gorm"
	"github.com/k-ueki/tmanager/server/domain/model"
	"github.com/k-ueki/tmanager/server/usecase"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type (
	FollowerController interface {
		Show(ctx echo.Context) error
		ListFollow(ctx echo.Context) error
		ListFollower(ctx echo.Context) error
		ListUnrequited(ctx echo.Context) error
		ListNew(ctx echo.Context) error
		ListBye(ctx echo.Context) error
		Hoge(ctx echo.Context) error
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
		return errors.Wrap(err, "failed to lis followers")
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
		return errors.Wrap(err, "failed to lis followers")
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
		return errors.Wrap(err, "failed to Show")
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
		return errors.Wrap(err, "failed to list unrequitedUsers")
	}
	return ctx.JSON(http.StatusOK, unrequitedUsers)
}

func (c *followerController) Hoge(ctx echo.Context) error {
	hoge, err := c.FollowerUseCase.GetFollowersIDsFromTwitterAPI()
	if err != nil {
		return errors.Wrap(err, "failed to lis followers")
	}
	return ctx.JSON(http.StatusOK, hoge)
}

func (c *followerController) ListNew(ctx echo.Context) error {
	user := model.User{
		ID:   1,
		Name: "User1",
	}

	newFollowers, err := c.FollowerUseCase.ListNewFollower(&user)
	if err != nil {
		return errors.Wrap(err, "failed to get new followers")
	}
	return ctx.JSON(http.StatusOK, newFollowers)
}

func (c *followerController) ListBye(ctx echo.Context) error {
	return nil
}
