package Entity

type MapPolicyImage struct {
	ID       int64 `json:"id"`
	PolicyID int64 `json:"policy-id"`
	ImageID  int64 `json:"image-id"`
}
