package JWT

import (
	"fmt"
	"github.com/Liang-jonas/jnote/Conf"
	"github.com/Liang-jonas/jnote/Models/DTO"
	"testing"
)

func TestJWT(t *testing.T) {
	c := Conf.JWT{
		Signature:       "Jonas",
		Issuer:          "Jonas",
		JwtId:           "Jonas",
		Subject:         "Jonas-note",
		LoginAccessTTL:  60,
		LoginRefreshTTL: 120,
	}
	engine := NewJwtEngine(c)
	u := &DTO.User{
		Username: "jonas",
		Group:    "test",
	}
	text, err := engine.GenAccess(u)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(text)
	ue, err := engine.Parse(text)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ue)
}
