package VO

type Note struct {
	Title      string `json:"title"`
	Ctx        string `json:"ctx"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}
