package entity

type (
	User struct {
		ID   int
		Name string
	}

	UserFollowerTid struct {
		UserID      uint
		FollowerTid uint64
	}
)

func (u *User) TableName() string {
	return "user"
}
