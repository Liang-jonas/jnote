package Redis

func (d *dbRepo) Close() error {
	d.cancel()
	return d.db.Close()
}
