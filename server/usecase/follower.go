package usecase

import (
	"github.com/jinzhu/gorm"
	"github.com/k-ueki/tmanager/server/domain/entity"
	"github.com/k-ueki/tmanager/server/domain/model"
	"github.com/k-ueki/tmanager/server/domain/repository"
	"github.com/k-ueki/tmanager/server/infra/presistence"
	"github.com/pkg/errors"
)

type (
	FollowerUseCase interface {
		Show(userID int) (*entity.User, error)
		ListFollower(user *model.User) ([]*entity.User, error)
		ListFollow(user *model.User) ([]*entity.User, error)
		ListUnrequitedUsers(user *model.User) ([]*entity.User, error)
	}

	followerUseCase struct {
		repository.FollowerRepository
	}
)

func NewFollowerUseCase(db *gorm.DB) FollowerUseCase {
	return &followerUseCase{presistence.NewFollowerPersistence(db)}
}

// The people following the User
func (u *followerUseCase) Show(userID int) (*entity.User, error) {
	user, err := u.FollowerRepository.FindByUserID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// The people following the User
func (u *followerUseCase) ListFollower(user *model.User) ([]*entity.User, error) {
	followers, err := u.FollowerRepository.FindFollowerByUserID(user.ID)
	if err != nil {
		return nil, err
	}
	return followers, nil
}

// The people the User follow
func (u *followerUseCase) ListFollow(user *model.User) ([]*entity.User, error) {
	followers, err := u.FollowerRepository.FindFollowByUserID(user.ID)
	if err != nil {
		return nil, err
	}
	return followers, nil
}

func (u *followerUseCase) ListUnrequitedUsers(user *model.User) ([]*entity.User, error) {
	unrequitedUsers, err := u.FollowerRepository.FindUnrequitedUser(user.ID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to list the people I follow")
	}

	return unrequitedUsers, nil
}

//var ucl = NewUsersClient()
//
//var mode string
//if r.ContentLength != 0 {
//mode = GetMode(r)
//}
//
//var dbh = &db.DBHandler{
//	DB: config.SetDB(),
//}
//
//pathToGetFollowers := baseURL + "followers/list.json"
//pathToGetIds := baseURL + "followers/ids.json"
//_, Ids := ucl.GetFollowersList(pathToGetFollowers, pathToGetIds)
////bodyF, Ids := ucl.GetFollowersList(pathToGetFollowers, pathToGetIds)
//
//if mode == "register" {
//_, fromdb := dbh.Select("followers")
//
////dbの情報とIdsを比較
//newf, byef := db.FindNewBye(&Ids, fromdb)
//fmt.Println("NEW", newf, "\nBYE", byef) //Ids
//
//type responseStruct struct {
//Mode  string       `json:mode`
//Users []users.User `json:users`
//}
//var resp = make([]responseStruct, 2)
//if len(byef.Ids) != 0 {
//resp[1].Mode = "bye"
//users := ucl.ConvertIdsToUsers(byef.Ids)
//resp[1].Users = users
//fmt.Println("RESP", resp)
//}
//
//if len(newf.Ids) != 0 {
//resp[0].Mode = "new"
//users := ucl.ConvertIdsToUsers(newf.Ids)
//resp[0].Users = users
//fmt.Println("RESP", resp)
//}
//
//bytes, _ := json.Marshal(&resp)
//fmt.Fprintf(w, string(bytes))
//
//return
//}
