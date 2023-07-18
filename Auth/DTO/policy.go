package DTO

import "github.com/Liang-jonas/golib/object"

type Policy struct {
	PID int64 `json:"pid"`
	UID int64 `json:"user-id"`
	GID int64 `json:"group-id"`
}

func (p *Policy) String() string {
	return object.ToString(p)
}
