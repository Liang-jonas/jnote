package Entity

type Policy struct {
	ID      int64 `json:"id"`
	UserID  int64 `json:"user-id"`
	GroupID int64 `json:"group-id"`
}
