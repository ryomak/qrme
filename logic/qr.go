package logic

import (
	"bytes"
	"image"
	"os"

	qrcode "github.com/skip2/go-qrcode"
)

func CreateQR(text string, size int) (image.Image, error) {
	host := os.Getenv("WEB_HOST")
	if host == "" {
		host = "http://xxxx.com"
	}
	var png []byte
	png, err := qrcode.Encode(text, qrcode.Medium, size)
	if err != nil {
		return nil, err
	}
	qrImage, _, err := image.Decode(bytes.NewReader(png))
	return qrImage, err
}
