package Redis

func (d *dbRepo) HGet(key, field string) (string, error) {
	res := d.db.HGet(d.dbKeyPrefix+key, field)
	return res.Val(), res.Err()
}
