package Redis

import (
	"github.com/Liang-jonas/jnote/Conf"
	"testing"
)

func TestRedis(t *testing.T) {
	a := Conf.Redis{
		Ip:             "192.168.1.108",
		Port:           "6379",
		Password:       "abcd1234",
		KeyPrefix:      "test",
		Dbname:         0,
		PoolSize:       1,
		ConnectTimeout: 500,
	}
	r, err := NewDB(a)
	if err != nil {
		t.Error(err)
	}

	/* if err := r.HSet("test", "test", "abcd"); err != nil {*/
	/*t.Error(err)*/
	/*}*/
	/*b, err := r.HExist("test", "tes")*/
	/*t.Error(b, err)*/
	/*s, err := r.HGet("test", "test")*/
	/*t.Error(s, err)*/
	/* if err := r.HDel("test", "test"); err != nil {*/
	/*t.Error(err)*/
	/*}*/
	/* c := []string{"a", "b", "c"}*/
	/*if err := r.SetSet("test", c); err != nil {*/
	/*t.Error(err)*/
	/*}*/
	/*t.Error(r.SetGetAll("test"))*/
	/*t.Error(r.SetExist("test", "a"))*/
	/*t.Error(r.SetExist("test", "d"))*/
	/*if err := r.SetSet("test", "e"); err != nil {*/
	/*t.Error(err)*/
	/*}*/
	t.Error(r.SGet("test"))
}
