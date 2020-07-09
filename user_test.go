package mysdk

import (
	"os"
	"testing"
)

func TestRongCloud_Register(t *testing.T) {
	rc := NewRongCloud(os.Getenv("APP_KEY"), os.Getenv("APP_SECRET"))
	resq, err := rc.Register("hejinyu", "hejnyu","hejinyu")
	if err  != nil {
		t.Errorf("%v", err)
	}
	t.Log(os.Getenv("APP_KEY"))
	t.Log(resq)
}
func TestRongCloud_Refresh(t *testing.T) {
	rc := NewRongCloud(os.Getenv("APP_KEY"), os.Getenv("APP_SECRET"))
	resq, err := rc.Refresh("hejinyu", "hejnyu","hejinyu")
	if err  != nil {
		t.Errorf("%v", err)
	}
	t.Log(os.Getenv("APP_KEY"))
	t.Log(resq)
}
func TestRongCloud_AddBlackList(t *testing.T) {
	rc := NewRongCloud(os.Getenv("APP_KEY"), os.Getenv("APP_SECRET"))
	resq, err := rc.AddBlackList("hejinyu", "hejnyu","hejinyu")
	if err  != nil {
		t.Errorf("%v", err)
	}
	t.Log(os.Getenv("APP_KEY"))
	t.Log(resq)
}
func TestRongCloud_RemoveBlackList(t *testing.T) {
	rc := NewRongCloud(os.Getenv("APP_KEY"), os.Getenv("APP_SECRET"))
	resq, err := rc.RemoveBlackList("hejinyu", "hejnyu","hejinyu")
	if err  != nil {
		t.Errorf("%v", err)
	}
	t.Log(os.Getenv("APP_KEY"))
	t.Log(resq)
}
func TestRongCloud_GetBlackList(t *testing.T) {
	rc := NewRongCloud(os.Getenv("APP_KEY"), os.Getenv("APP_SECRET"))
	resq, err := rc.GetBlackList("hejinyu")
	if err  != nil {
		t.Errorf("%v", err)
	}
	t.Log(os.Getenv("APP_KEY"))
	t.Log(resq)
}
func TestRongCloud_AddBlock(t *testing.T) {
	rc := NewRongCloud(os.Getenv("APP_KEY"), os.Getenv("APP_SECRET"))
	resq, err := rc.AddBlock(100, "hejnyu","hejinyu")
	if err  != nil {
		t.Errorf("%v", err)
	}
	t.Log(os.Getenv("APP_KEY"))
	t.Log(resq)
}
func TestRongCloud_BlockList(t *testing.T) {
	rc := NewRongCloud(os.Getenv("APP_KEY"), os.Getenv("APP_SECRET"))
	resq, err := rc.BlockList(1, 10)
	if err  != nil {
		t.Errorf("%v", err)
	}
	t.Log(os.Getenv("APP_KEY"))
	t.Log(resq)
}
func TestRongCloud_RemoveBlock(t *testing.T) {
	rc := NewRongCloud(os.Getenv("APP_KEY"), os.Getenv("APP_SECRET"))
	resq, err := rc.RemoveBlock("hejinyu", "hejnyu","hejinyu")
	if err  != nil {
		t.Errorf("%v", err)
	}
	t.Log(os.Getenv("APP_KEY"))
	t.Log(resq)
}
func TestRongCloud_AddWhiteList(t *testing.T) {
	rc := NewRongCloud(os.Getenv("APP_KEY"), os.Getenv("APP_SECRET"))
	resq, err := rc.AddWhiteList("hejinyu", "hejnyu","hejinyu")
	if err  != nil {
		t.Errorf("%v", err)
	}
	t.Log(os.Getenv("APP_KEY"))
	t.Log(resq)
}
func TestRongCloud_RemoveWhiteList(t *testing.T) {
	rc := NewRongCloud(os.Getenv("APP_KEY"), os.Getenv("APP_SECRET"))
	resq, err := rc.RemoveWhiteList("hejinyu", "hejnyu","hejinyu")
	if err  != nil {
		t.Errorf("%v", err)
	}
	t.Log(os.Getenv("APP_KEY"))
	t.Log(resq)
}
func TestRongCloud_GetWhiteList(t *testing.T) {
	rc := NewRongCloud(os.Getenv("APP_KEY"), os.Getenv("APP_SECRET"))
	resq, err := rc.GetWhiteList("hejinyu")
	if err  != nil {
		t.Errorf("%v", err)
	}
	t.Log(os.Getenv("APP_KEY"))
	t.Log(resq)
}
func TestRongCloud_TagSet(t *testing.T) {
	rc := NewRongCloud(os.Getenv("APP_KEY"), os.Getenv("APP_SECRET"))
	resq, err := rc.TagSet("hejinyu", "hejnyu","hejinyu")
	if err  != nil {
		t.Errorf("%v", err)
	}
	t.Log(os.Getenv("APP_KEY"))
	t.Log(resq)
}
func TestRongCloud_TagGet(t *testing.T) {
	rc := NewRongCloud(os.Getenv("APP_KEY"), os.Getenv("APP_SECRET"))
	resq, err := rc.TagGet("hejinyu")
	if err  != nil {
		t.Errorf("%v", err)
	}
	t.Log(os.Getenv("APP_KEY"))
	t.Log(resq)
}
func TestRongCloud_TagBatchSet(t *testing.T) {
	rc := NewRongCloud(os.Getenv("APP_KEY"), os.Getenv("APP_SECRET"))
	resq, err := rc.TagBatchSet([]string{"hejinyu","hejinyu1"}, []string{"a", "b"})
	if err  != nil {
		t.Errorf("%v", err)
	}
	t.Log(os.Getenv("APP_KEY"))
	t.Log(resq)
}
