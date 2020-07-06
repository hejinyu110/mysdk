package mysdk

import (
	"os"
	"testing"
)

func TestRongCloud_Register(t *testing.T) {
	os.Setenv("APP_KEY","mgb7ka1nmeqig")
	os.Setenv("APP_KEY","g3ScjDkuHanSCg")
	rc := NewRongCloud(os.Getenv("APP_KEY"), os.Getenv("APP_SECRET"))
	resq, err := rc.Register("hejinyu", "hejnyu","hejinyu")
	if err  != nil {
		t.Errorf("%v", err)
	}
	t.Log(os.Getenv("APP_KEY"))
	t.Log(resq)
}
