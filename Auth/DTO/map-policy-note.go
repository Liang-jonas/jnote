package DTO

import "github.com/Liang-jonas/golib/object"

type MapPolicyNote struct {
	ID       int64 `json:"id"`
	NoteID   int64 `json:"note-id"`
	PolicyID int64 `json:"policy-id"`
}

func (m *MapPolicyNote) String() string {
	return object.ToString(m)
}
