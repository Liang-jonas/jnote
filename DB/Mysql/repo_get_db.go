package Mysql

import "database/sql"

func (d *dbRepo) GetDB() *sql.DB {
	return d.db
}
