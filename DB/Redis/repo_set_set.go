package Redis

func (d *dbRepo) SetSet(key string, value ...interface{}) error {
	return d.db.SAdd(d.dbKeyPrefix+key, value...).Err()
}
