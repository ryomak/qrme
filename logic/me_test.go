package logic

import (
	"bytes"
	"image"
	_ "image/jpeg"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

func TestCreateImageByBackgroundImage(t *testing.T) {
	res, err := http.Get("https://ryomak.github.io/dollarphin/static/img/about-us.284937b.jpg")
	if err != nil {
		t.Errorf("Create() error = %v", err)
		return
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Create() error = %v", err)
		return
	}
	backgroundImage, _, err := image.Decode(bytes.NewBuffer(data))
	if err != nil || res.StatusCode != 200 {
		t.Errorf("create() error = %v, statuCode = %v", err, res.StatusCode)
		return
	}
	me, err := NewMeBuilder().NickName("test").Introduction("暇です").Setting(nil, nil, &backgroundImage).Build()
	if err != nil {
		t.Errorf("Create() error = %v", err)
		return
	}
	file, err := os.Create("testCreateImageByBackgroundImage.png")
	if err != nil {
		t.Errorf("Create() error = %v", err)
		return
	}
	defer file.Close()
	if err := me.CreateImage(file); err != nil {
		t.Errorf("Create() error = %v", err)
		return
	}
}

func TestCreateImage(t *testing.T) {
	me, err := NewMeBuilder().NickName("test").Introduction("暇です").Build()
	if err != nil {
		t.Errorf("Create() error = %v", err)
		return
	}
	file, err := os.Create("testCreateImage.png")
	if err != nil {
		t.Errorf("Create() error = %v", err)
		return
	}
	defer file.Close()
	if err := me.CreateImage(file); err != nil {
		t.Errorf("Create() error = %v", err)
		return
	}
}
