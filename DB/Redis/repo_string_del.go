package Redis

func (d *dbRepo) SDel(key string) error {
	return d.db.Del(d.dbKeyPrefix + key).Err()
}
