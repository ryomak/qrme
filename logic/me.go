package logic

import (
	"image"
	"image/draw"
	"image/png"
	"io"

	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

type Me struct {
	NickName     string
	SubName      *string
	Introduction string
	Icon         string
	Languages    []Language
	Socials      []Social
	Setting      *Setting
}

func (me *Me) CreateImage(w io.Writer) error {
	img := image.NewRGBA(image.Rect(0, 0, me.Setting.Width, me.Setting.Height))

	if me.Setting.HasBackgroundImage() {
		draw.Draw(img, img.Bounds(), *me.Setting.BackgroundImage, image.ZP, draw.Src)
	} else {
		draw.Draw(img, img.Bounds(), me.Setting.BackgroundColor.ToUniform(), image.ZP, draw.Src)
	}
	dr := &font.Drawer{
		Dst:  img,
		Src:  me.Setting.Color.ToUniform(),
		Face: me.Setting.FontFace,
		Dot:  fixed.Point26_6{},
	}
	dr.Dot.X = fixed.I(10)
	dr.Dot.Y = fixed.I(100)
	dr.DrawString(me.NickName)
	return png.Encode(w, img)
}

func SplitByMeasureWidth(text string, maxWidth int, dr *font.Drawer) []string {
	var (
		lines []string
		line  string
	)
	for _, v := range text {
		vs := string(v)
		w := dr.MeasureString(line + vs).Round()
		switch {
		case maxWidth <= w:
			lines = append(lines, line)
			line = vs
		default:
			line = line + vs
		}
	}
	lines = append(lines, line)
	return lines
}
