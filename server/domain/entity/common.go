package entity

type (
	TwitterID struct {
		ID uint64 `json:"id"`
	}
)

func ConvertTwitterIDToUint64(ids []*TwitterID) []*uint64 {
	var resp []*uint64
	for _, id := range ids {
		resp = append(resp, &id.ID)
	}
	return resp
}
