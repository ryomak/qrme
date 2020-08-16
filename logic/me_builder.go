package logic

import (
	"image"
	"log"
)

type MeBuilder struct {
	me *Me
}

func NewMeBuilder() *MeBuilder {
	return &MeBuilder{
		me: &Me{},
	}
}

func (mb *MeBuilder) NickName(nickname string) *MeBuilder {
	mb.me.NickName = nickname
	return mb
}

func (mb *MeBuilder) SubName(name string) *MeBuilder {
	mb.me.SubName = &name
	return mb
}

func (mb *MeBuilder) Introduction(introduction string) *MeBuilder {
	mb.me.Introduction = introduction
	return mb
}

func (mb *MeBuilder) Setting(color, backgroundColor *Color, backgroundImage *image.Image) *MeBuilder {
	if color == nil {
		c := defaultColor["color"]
		color = &c
	}
	if backgroundColor == nil {
		c := defaultColor["background"]
		backgroundColor = &c
	}
	mb.me.Setting = &Setting{
		Color:           *color,
		BackgroundColor: backgroundColor,
		BackgroundImage: backgroundImage,
		Width:           1200,
		Height:          630,
	}
	if err := mb.me.Setting.SetFontFromFile("../font/Koruri-Light.ttf"); err != nil {
		log.Fatal(err)
	}

	return mb
}

func (mb *MeBuilder) Icon(icon string) *MeBuilder {
	mb.me.Icon = icon
	return mb
}

func (mb *MeBuilder) Languages(languages []string) *MeBuilder {
	var langs []Language
	for _, v := range languages {
		l, exist := languageMap[v]
		if exist {
			langs = append(langs, l)
		}
	}
	mb.me.Languages = langs
	return mb
}

func (mb *MeBuilder) Social(socials []Social) *MeBuilder {
	var ss []Social
	for _, v := range socials {
		url, exist := socialMap[v.Name]
		if exist {
			v.URL = url
			ss = append(ss, v)
		}
	}
	mb.me.Socials = ss
	return mb
}

func (mb *MeBuilder) Build() *Me {
	if mb.me.Setting == nil {
		mb.Setting(nil, nil, nil)
	}
	if mb.me.Languages == nil {
		mb.me.Languages = []Language{}
	}
	if mb.me.Socials == nil {
		mb.me.Socials = []Social{}
	}
	return mb.me
}
