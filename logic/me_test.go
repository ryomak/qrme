package logic

import (
	"os"
	"testing"
)

func TestCreateImage(t *testing.T) {
	me := NewMeBuilder().NickName("test").Introduction("暇です").Build()
	file, err := os.Create("test.png")
	if err != nil {
		t.Errorf("Create() error = %v", err)
		return
	}
	if err := me.CreateImage(file); err != nil {
		t.Errorf("Create() error = %v", err)
		return
	}
	file.Close()
}
