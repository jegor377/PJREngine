package GameEngine

import "github.com/veandco/go-sdl2/sdl"

type ButtonManager struct {
	buttons []*Button
	pos sdl.Point
	margin int32
	buttonHeight int32
}

func NewButtonManager(heightOffset int32) *ButtonManager {
	return &ButtonManager{
		make([]*Button, 0),
		sdl.Point{400, heightOffset},
		10,
		50,
	}
}

func (btnMgr *ButtonManager) SetHeightOffset(offset int32) {
	btnMgr.pos.Y = offset
}

func (btnMgr *ButtonManager) Add(text string, target uint32, source uint32, color uint8, fontMgr *FontManager) {
	btnOffset := ( int32(len(btnMgr.buttons) + 1) * btnMgr.buttonHeight )
	btnMarginOffset := (int32(len(btnMgr.buttons) + 1) * btnMgr.margin)
	YaxisPos := btnMgr.pos.Y + btnOffset + btnMarginOffset
	dst := sdl.Rect{400, YaxisPos, 360, btnMgr.buttonHeight}
	btnMgr.buttons = append(btnMgr.buttons, NewButton(text, target, source, color, dst, fontMgr))
}

func (btnMgr *ButtonManager) Clear() {
	btnMgr.buttons = nil
	btnMgr.buttons = make([]*Button, 0)
}

func (btnMgr *ButtonManager) Destroy() {
	btnMgr.Clear()
}

func (btnMgr *ButtonManager) GetButtons() []*Button {
	return btnMgr.buttons
}
