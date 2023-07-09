package Redis

func (d *dbRepo) HSet(key, field string, value interface{}) error {
	return d.db.HSet(d.dbKeyPrefix+key, field, value).Err()
}
