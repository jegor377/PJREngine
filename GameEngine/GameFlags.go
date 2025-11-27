package GameEngine

import "github.com/veandco/go-sdl2/sdl"

type GameFlags struct {
	ScreenMode uint32
}

func NewGameFlags() *GameFlags {
	return &GameFlags{
		sdl.WINDOW_SHOWN,
	}
}
