package Redis

func (d *dbRepo) SGet(key string) (string, error) {
	res := d.db.Get(d.dbKeyPrefix + key)
	return res.Val(), res.Err()
}
