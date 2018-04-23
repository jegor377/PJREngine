package GameEngine

import "github.com/veandco/go-sdl2/sdl"
import "fmt"

func (engine *GameEngine) loadTextures() error {
	if !engine.sprites.IsEmpty {
		fmt.Println("Loading sprites...")
		size := float32( len(engine.sprites.Sprites) )
		var actualState float32 = 0
		for id, v := range engine.sprites.Sprites {
			if err := engine.textureMgr.LoadTexture(engine.renderer, uint32(id), v); err != nil {
				return err
			}
			actualState++
			fmt.Print("Loading sprites [", int( (actualState/size) * 100.0 ), "%]\n")
		}
	}
	return nil
}

func (engine *GameEngine) changeActualSprite(spriteId uint32) error {
	var err error
	if engine.actualSprite, err = engine.textureMgr.GetTexture(spriteId); err != nil {
		return err
	}
	return nil
}

// returns src, dst, err
func (engine *GameEngine) getSpriteRect() (sdl.Rect, sdl.Rect, error) {
	_, _, w, h, err := engine.actualSprite.Query()
	if err != nil {
		return sdl.Rect{0, 0, 0, 0}, sdl.Rect{0, 0, 0, 0}, err
	}
	return sdl.Rect{0, 0, w, h}, sdl.Rect{30, 30, 340, 540}, nil
}