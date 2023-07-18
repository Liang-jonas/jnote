package DTO

import "github.com/Liang-jonas/golib/object"

type MapPolicyImage struct {
	ID       int64 `json:"id"`
	ImageID  int64 `json:"image-id"`
	PolicyID int64 `json:"policy-id"`
}

func (m *MapPolicyImage) String() string {
	return object.ToString(m)
}
