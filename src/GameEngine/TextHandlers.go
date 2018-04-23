package GameEngine

import "github.com/veandco/go-sdl2/sdl"
import "Informations"
import "errors"
import "strings"

func (engine *GameEngine) changeActualText(info *Informations.DialogInformations) error {
	preparedText := strings.Replace(info.Text, "\\n", "\n", -1) + "\n"
	// if !engine.appState.IsEnding {
	// 	wydt, err := engine.getWYDT()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	preparedText += wydt
	// }
	texture, err := engine.fontMgr.GetTexture(engine.renderer, "default", preparedText, getColorById(info.Color))
	if err != nil {
		panic(err)
	}
	engine.actualText.Destroy()
	engine.actualText = texture
	_, offset := engine.getTextLength()
	engine.buttonMgr.SetHeightOffset(offset)
	return nil
}

//src, dst, err
func (engine *GameEngine) getTextRect() (sdl.Rect, sdl.Rect, error) {
	_, _, w, h, err := engine.actualText.Query()
	if err != nil {
		return sdl.Rect{0, 0, 0, 0}, sdl.Rect{0, 0, 0, 0}, err
	}
	if h > MAX_TEXT_SIZE {
		return sdl.Rect{0, engine.textMouseOffset * -1, w, MAX_TEXT_SIZE}, sdl.Rect{400, 35, w, MAX_TEXT_SIZE}, nil
	}
	return sdl.Rect{0, 0, w, h}, sdl.Rect{400, 35, w, h}, nil
}

func (engine *GameEngine) getTextLength() (w, h int32) {
	_, _, w, h, err := engine.actualText.Query()
	if err != nil {
		panic(err)
	}
	if h > 230 {
		return w, 230
	}
	return w, h
}

func (engine *GameEngine) getActualTextLength() (w, h int32) {
	_, _, w, h, err := engine.actualText.Query()
	if err != nil {
		panic(err)
	}
	return w, h
}

func (engine *GameEngine) getWYDT() (string, error) {
	if wydt, ok := engine.appState.WhatYouDoText.Get().(string); ok {
		return wydt, nil
	}
	return "", errors.New("FATAL ERROR: CANNOT TO LOAD WYDT. CONTACT DEVELOPER - ERROR W CHUJ. SKONTAKTUJ SIE Z IGOREM. | TextHandlers.go")
}