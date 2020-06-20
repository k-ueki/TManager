package model

type (
	User struct {
		ID   int
		Name string
	}

	UserFromTwitterAPI struct {
		Users []UserInfoFromTwitterAPI `json:"users"`
	}

	UserInfoFromTwitterAPI struct {
		ID          int64  `json:"id"`
		IDStr       string `json:"id_str"`
		Name        string `json:"name"`
		ScreenName  string `json:"screen_name"`
		Description string `json:"description"`
		URL         string `json:"url"`
	}

	UserIDs struct {
		IDs []int64 `json:"ids"`
	}
)

func (ui *UserIDs) ConvertUserIDsToUint64() []*uint64 {
	var resp []*uint64
	for _, id := range ui.IDs {
		tmp := uint64(id)
		resp = append(resp, &tmp)
	}
	return resp
}
