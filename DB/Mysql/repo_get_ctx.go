package Mysql

import "context"

func (d *dbRepo) GetCtx() context.Context {
	return d.ctx
}
