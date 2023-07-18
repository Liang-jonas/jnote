package VO

type NoteListItem struct {
	ID     int64          `json:"id"`
	Title  string         `json:"title"`
	Type   bool           `json:"type"`
	SubDir []NoteListItem `json:"subDir"`
}
