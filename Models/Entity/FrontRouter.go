package Entity

type FrontRout struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Path        string `json:"path"`
	Description string `json:"description"`
	Type        bool   `json:"type"`
	Index       int    `json:"index"`
}
