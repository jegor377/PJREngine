package Sections

import "bytes"

type SpritesSection struct {
	Sprites []string
	IsEmpty bool
}

func NewSpritesSection() *SpritesSection {
	return &SpritesSection{
		make([]string, 0),
		false,
	}
}

func (sprites *SpritesSection) DecodeSprites(byteCode []byte) {
	if len(byteCode) == 0 {
		sprites.IsEmpty = true
		return
	}
	sprites.Sprites = append(sprites.Sprites, "none") // first must be none
	paths := bytes.Split(byteCode, []byte{0})
	for _, e := range(paths) {
		if len(e) > 0 {
			sprites.Sprites = append(sprites.Sprites, string(e))
		}
	}
}