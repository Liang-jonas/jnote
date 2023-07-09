package Entity

type MapPolicyNote struct {
	ID       int64 `json:"id"`
	PolicyID int64 `json:"policy-id"`
	NoteID   int64 `json:"note-id"`
}
