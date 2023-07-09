package JWT

func (e *engine) GetRefreshTTL() int {
	return int(e.LoginRefreshTTL)
}
