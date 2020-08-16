package logic

import (
	"image"
	"image/color"
	"io/ioutil"
	"path/filepath"

	_ "github.com/ryomak/qrme/logic/statik"

	"github.com/rakyll/statik/fs"
	"golang.org/x/image/font"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

type Setting struct {
	Color           Color
	BackgroundColor *Color
	BackgroundImage *image.Image
	FontFace        map[string]font.Face
	Width           int
	Height          int
}

func (s Setting) HasBackgroundImage() bool {
	return s.BackgroundImage != nil
}

func (s Setting) GetBigFontSize() float64 {
	return float64(s.Height) / 5
}

func (s Setting) GetMidFontSize() float64 {
	return float64(s.Height) / 8
}

func (s Setting) GetSmallFontSize() float64 {
	return float64(s.Height) / 10
}

func (s *Setting) SetFontFromFile(filename string) error {
	statikFS, err := fs.New()
	if err != nil {
		return err
	}
	r, err := statikFS.Open(filepath.Join("/", "font", filename))
	if err != nil {
		return err
	}
	defer r.Close()
	fntBytes, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	fnt, err := freetype.ParseFont(fntBytes)
	if err != nil {
		return err
	}
	faceMap := make(map[string]font.Face)
	faceMap["big"] = truetype.NewFace(fnt, &truetype.Options{
		Size: s.GetBigFontSize(),
	})
	faceMap["mid"] = truetype.NewFace(fnt, &truetype.Options{
		Size: s.GetMidFontSize(),
	})
	faceMap["small"] = truetype.NewFace(fnt, &truetype.Options{
		Size: s.GetSmallFontSize(),
	})
	s.FontFace = faceMap
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
