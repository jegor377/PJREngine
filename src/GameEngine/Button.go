package GameEngine

import "github.com/veandco/go-sdl2/sdl"

type Button struct {
	target uint32
	source uint32
	text string
	color uint8
	dst sdl.Rect
	fontMgr *FontManager
}

func NewButton(text string, target uint32, source uint32, color uint8, dst sdl.Rect, fontMgr *FontManager) *Button {
	return &Button{
		target,
		source,
		text,
		color,
		dst,
		fontMgr,
	}
}

func (btn *Button) HasBeenClicked(x, y int32, which, state uint8) bool {
	return (int32(x) > btn.dst.X) && (int32(x) < btn.dst.X + btn.dst.W) && (int32(y) > btn.dst.Y) && (int32(y) < btn.dst.Y + btn.dst.H) && (state == 1) && (which == 1)
}

func (btn *Button) GetTarget() uint32 {
	return btn.target
}

func (btn *Button) GetSource() uint32 {
	return btn.source
}

func (btn *Button) GetTexture(renderer *sdl.Renderer) (*sdl.Rect, *sdl.Texture, *sdl.Rect) {
	texture, err := btn.fontMgr.GetTexture(renderer, "option", btn.text, getColorById(btn.color))
	if err != nil {
		panic(err)
	}
	_, _, w, h, err := texture.Query()
	if err != nil {
		panic(err)
	}
	pos := sdl.Rect{btn.dst.X + (btn.dst.W / 2) - (w/2), btn.dst.Y + (btn.dst.H / 2) - (h/2), w, h}
	return &btn.dst, texture, &pos
}