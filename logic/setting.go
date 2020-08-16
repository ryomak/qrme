package logic

import (
	"image"
	"image/color"
	"io/ioutil"

	"golang.org/x/image/font"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

type Setting struct {
	Color           Color
	BackgroundColor *Color
	BackgroundImage *image.Image
	FontFace        font.Face
	Width           int
	Height          int
}

func (s Setting) HasBackgroundImage() bool {
	return s.BackgroundImage != nil
}

func (s *Setting) SetFontFromFile(filepath string) error {
	opt := truetype.Options{
		Size: FontSize,
	}
	fntBytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}
	fnt, err := freetype.ParseFont(fntBytes)
	if err != nil {
		return err
	}
	face := truetype.NewFace(fnt, &opt)
	s.FontFace = face
	return nil
}

type Color struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

func (c Color) ToUniform() *image.Uniform {
	col := color.RGBA{c.R, c.G, c.B, c.A}
	return image.NewUniform(col)
}

var defaultColor = map[string]Color{
	"color": {
		R: 20,
		G: 30,
		B: 30,
		A: 255,
	},
	"background": {
		R: 255,
		G: 255,
		B: 255,
		A: 255,
	},
}

const (
	FontSize       = 84.0
	FontLineHeight = 1.3
)
