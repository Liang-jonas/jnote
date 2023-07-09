package Conf

import (
	"testing"
)

func TestCon(t *testing.T) {
	c := NewConf()
	err := c.Read("../config.ini")
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	t.Logf("%s", c.GetCfg())
}
