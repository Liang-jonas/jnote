package Redis

import "time"

func (d *dbRepo) SSet(key, value string, ttl int) error {
	return d.db.Set(d.dbKeyPrefix+key, value, time.Duration(ttl)*time.Second).Err()
}
