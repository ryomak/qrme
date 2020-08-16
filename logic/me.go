package logic

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"io"
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

type Me struct {
	Unique       string
	NickName     string
	SubName      *string
	Introduction string
	Icon         string
	Languages    []Language
	Socials      []Social
	Setting      *Setting
}

func (me Me) GetWebURL() string {
	host := os.Getenv("WEB_HOST")
	if host == "" {
		host = "http://xxxx.com"
	}
	return fmt.Sprintf("%s/profile/%s", host, me.Unique)
}

func (me Me) PostWebURL() string {
	host := os.Getenv("WEB_HOST")
	if host == "" {
		host = "http://xxxx.com"
	}
	return fmt.Sprintf("%s/profile", host)
}

func (me *Me) CreateImage(w io.Writer) error {
	img := image.NewRGBA(image.Rect(0, 0, me.Setting.Width, me.Setting.Height))

	// write background
	if me.Setting.HasBackgroundImage() {
		backgroundImage := *me.Setting.BackgroundImage
		draw.Draw(img, backgroundImage.Bounds(), backgroundImage, image.ZP, draw.Src)
	} else {
		draw.Draw(img, img.Bounds(), me.Setting.BackgroundColor.ToUniform(), image.ZP, draw.Src)
	}

	dr := &font.Drawer{
		Dst:  img,
		Src:  me.Setting.Color.ToUniform(),
		Face: me.Setting.FontFace["big"],
		Dot:  fixed.Point26_6{},
	}
	dr.Dot.X = fixed.I(10)
	dr.Dot.Y = fixed.I(100)
	dr.DrawString(me.NickName)

	qr, err := CreateQR(me.GetWebURL(), 256)
	if err != nil {
		return err
	}
	draw.Draw(img, image.Rect(300, 300, 300+256, 300+256), qr, image.ZP, draw.Src)
	return png.Encode(w, img)
}
