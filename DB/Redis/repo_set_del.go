package Redis

func (d *dbRepo) SetDel(key, field string) error {
	return d.db.SRem(d.dbKeyPrefix+key, field).Err()
}
