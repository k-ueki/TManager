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
	FollowerQueryController interface {
		Show(ctx echo.Context) error
		ListFollow(ctx echo.Context) error
		ListFollower(ctx echo.Context) error
		ListUnrequited(ctx echo.Context) error
		ListNew(ctx echo.Context) error
		ListBye(ctx echo.Context) error
		Hoge(ctx echo.Context) error
	}

	followerQueryController struct {
		usecase.FollowerQueryUseCase
	}
)

func NewFollowerController(db *gorm.DB) FollowerQueryController {
	return &followerQueryController{usecase.NewFollowerUseCase(db)}
}

func (c *followerQueryController) ListFollow(ctx echo.Context) error {
	user := model.User{
		ID:   1,
		Name: "User1",
	}

	followers, err := c.FollowerQueryUseCase.ListFollow(&user)
	if err != nil {
		return errors.Wrap(err, "failed to lis followers")
	}
	return ctx.JSON(http.StatusOK, followers)
}

func (c *followerQueryController) ListFollower(ctx echo.Context) error {
	user := model.User{
		ID:   1,
		Name: "User1",
	}

	followers, err := c.FollowerQueryUseCase.ListFollower(&user)
	if err != nil {
		return errors.Wrap(err, "failed to lis followers")
	}
	return ctx.JSON(http.StatusOK, followers)
}

func (c *followerQueryController) Show(ctx echo.Context) error {
	param := input.ShowFollowerParam{}
	if err := ctx.Bind(&param); err != nil {
		return ctx.JSON(http.StatusBadRequest, "invalid params")
	}

	followers, err := c.FollowerQueryUseCase.Show(param.UserID)
	if err != nil {
		return errors.Wrap(err, "failed to Show")
	}
	return ctx.JSON(http.StatusOK, followers)
}

func (c *followerQueryController) ListUnrequited(ctx echo.Context) error {
	user := model.User{
		ID:   1,
		Name: "User1",
	}

	unrequitedUsers, err := c.FollowerQueryUseCase.ListUnrequitedUsers(&user)
	if err != nil {
		return errors.Wrap(err, "failed to list unrequitedUsers")
	}
	return ctx.JSON(http.StatusOK, unrequitedUsers)
}

func (c *followerQueryController) Hoge(ctx echo.Context) error {
	hoge, err := c.FollowerQueryUseCase.GetFollowersIDsFromTwitterAPI()
	if err != nil {
		return errors.Wrap(err, "failed to lis followers")
	}
	return ctx.JSON(http.StatusOK, hoge)
}

func (c *followerQueryController) ListNew(ctx echo.Context) error {
	user := model.User{
		ID:   1,
		Name: "User1",
	}

	newFollowers, err := c.FollowerQueryUseCase.ListNewFollower(&user)
	if err != nil {
		return errors.Wrap(err, "failed to get new followers")
	}
	return ctx.JSON(http.StatusOK, newFollowers)
}

func (c *followerQueryController) ListBye(ctx echo.Context) error {
	return nil
}
