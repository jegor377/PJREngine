package GameEngine

import "github.com/veandco/go-sdl2/sdl"

func (engine *GameEngine) Loop() error {
	engine.loadState()
	for !engine.quit {
		if err := engine.handleEvents(); err != nil {
			return err
		}
		if err := engine.update(); err != nil {
			return err
		}
		if err := engine.render(); err != nil {
			return err
		}
	}
	return nil
}

func (engine *GameEngine) handleEvents() error {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch e := event.(type) {
		case *sdl.QuitEvent:
			engine.exit()
			break
		case *sdl.MouseButtonEvent:
			if !engine.appState.IsEnding {
				engine.checkEventsForButtons(e)
			}
			break
		case *sdl.KeyDownEvent:
			switch e.Keysym.Sym {
			case sdl.K_ESCAPE:
				engine.exit()
				break
			case sdl.K_DOWN:
				engine.scrollText(-1)
				break
			case sdl.K_UP:
				engine.scrollText(1)
				break
			}
			break
		case *sdl.MouseWheelEvent:
			engine.scrollText(e.Y)
			break
		}
	}
	return nil
}

func (engine *GameEngine) scrollText(direction int32) {
	_, h := engine.getActualTextLength()
	newOffset := engine.textMouseOffset + direction * 5
	if (newOffset > (h - MAX_TEXT_SIZE) * -1) && (newOffset <= 0) {
		engine.textMouseOffset = newOffset
	}
}

func (engine *GameEngine) exit() {
	engine.quit = true
	engine.saveState()
}

func (engine *GameEngine) update() error {
	if err := engine.WatchSettings(); err != nil {
		return err
	}
	err := engine.checkForActualDialog()
	return err
}

func (engine *GameEngine) render() error {
	engine.renderer.Clear()
	if err := engine.printSprite(); err != nil {
		return err
	}
	if err := engine.printText(); err != nil {
		return err
	}
	for _, v := range engine.buttonMgr.GetButtons() {
		btnDST, texture, btnTextDST := v.GetTexture(engine.renderer)
		engine.renderer.Copy(engine.btnSprite, nil, btnDST)
		engine.renderer.Copy(texture, nil, btnTextDST)
		texture.Destroy()
		btnDST = nil
		btnTextDST = nil
	}
	engine.renderer.Present()
	return nil
}

func (engine *GameEngine) printSprite() error {
	if engine.actualSprite != nil {
		src, dst, err := engine.getSpriteRect()
		if err != nil {
			return err
		}
		engine.renderer.Copy(engine.actualSprite, &src, &dst)
	}
	return nil
}

func (engine *GameEngine) printText() error {
	if engine.actualText != nil {
		src, dst, err := engine.getTextRect()
		if err != nil {
			return err
		}
		engine.renderer.Copy(engine.actualText, &src, &dst)
	}
	return nil
}

func (engine *GameEngine) WatchSettings() error {
	if err := engine.checkForAppName(); err != nil {
		return err
	}
	if err := engine.checkForBgColor(); err != nil {
		return err
	}
	err := engine.checkForBgSong()
	return err
}