package JWT

func (e *engine) GetAccessTTL() int {
	return int(e.LoginAccessTTL)
}
