package src

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"os"

	"github.com/ryomak/qrme/logic"
)

func Create(me *logic.Me) error {
	file, err := os.Create("me.png")
	if err != nil {
		return err
	}
	defer file.Close()
	if err := me.CreateImage(file); err != nil {
		return err
	}
	// image to save gcs
	fi, _ := file.Stat()
	size := fi.Size()
	data := make([]byte, size)
	if _, err := file.Read(data); err != nil {
		return err
	}
	var params struct {
		Uid   string `json:"uid"`
		Image string `json:"image"`
	}
	params.Uid = me.Unique
	params.Image = base64.StdEncoding.EncodeToString(data)

	jsonData, err := json.Marshal(params)
	if err != nil {
		return err
	}
	if _, err := http.Post(me.PostWebURL(), "application/json", bytes.NewBuffer(jsonData)); err != nil {
		return err
	}
	return nil
}
