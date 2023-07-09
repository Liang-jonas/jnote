package Redis

func (d *dbRepo) HExist(key string, field string) (bool, error) {
	return d.db.HExists(d.dbKeyPrefix+key, field).Result()
}
