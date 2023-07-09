package Redis

func (d *dbRepo) SetGetAll(key string) ([]string, error) {
	return d.db.SMembers(d.dbKeyPrefix + key).Result()
}
