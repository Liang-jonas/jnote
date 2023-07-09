package Redis

func (d *dbRepo) Del(key string) error {
	return d.db.Del(d.dbKeyPrefix + key).Err()
}
