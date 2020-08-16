package src

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
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

	buf := bytes.NewBuffer(nil)
	writer := io.MultiWriter(file, buf)
	if err := me.CreateImage(writer); err != nil {
		return err
	}
	var params struct {
		Uid   string `json:"uid"`
		Image string `json:"image"`
	}
	params.Uid = me.Unique
	params.Image = base64.StdEncoding.EncodeToString(buf.Bytes())

	jsonData, err := json.Marshal(params)
	if err != nil {
		return err
	}
	if _, err := http.Post(me.PostWebURL(), "application/json", bytes.NewBuffer(jsonData)); err != nil {
		return err
	}
	return nil
}
