package Redis

func (d *dbRepo) SLen(key string) (int, error) {
	res, err := d.db.StrLen(d.dbKeyPrefix + key).Result()
	return int(res), err
}
