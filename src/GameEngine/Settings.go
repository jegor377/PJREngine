package GameEngine

import "github.com/veandco/go-sdl2/img"

func (engine *GameEngine) checkForAppName() error {
	if engine.appState.AppName.HasChanged() {
		val, _ := engine.appState.AppName.Get().(string)
		engine.window.SetTitle(val)
	}
	return nil
}

func (engine *GameEngine) checkForBgColor() error {
	if engine.appState.BgColor.HasChanged() {
		val, _ := engine.appState.BgColor.Get().(uint8)
		c := getColorById(val)
		err := engine.renderer.SetDrawColor(c.R, c.G, c.B, c.A)
		return err
	}
	return nil
}

func (engine *GameEngine) checkForBgSong() error {
	if engine.appState.BgSong.HasChanged() {
		val, _ := engine.appState.BgSong.Get().(uint32)
		if err := engine.chunkMgr.PlayRepeated(val, -1); err != nil {
			panic(err)
		}
	}
	return nil
}

func (engine *GameEngine) loadBtnSprite() {
	texture, err := img.LoadTexture(engine.renderer, "assets/DialogBtn.png")
	if err != nil {
		panic(err)
	}
	engine.btnSprite = texture
}