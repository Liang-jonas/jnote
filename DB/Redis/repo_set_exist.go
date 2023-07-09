package Redis

func (d *dbRepo) SetExist(key, value string) (bool, error) {
	return d.db.SIsMember(d.dbKeyPrefix+key, value).Result()
}
