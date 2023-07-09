package Redis

func (d *dbRepo) HDel(key, field string) error {
	return d.db.HDel(d.dbKeyPrefix+key, field).Err()
}
