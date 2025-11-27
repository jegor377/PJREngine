package main

import "fmt"
import "os"
import "github.com/veandco/go-sdl2/sdl"
import "github.com/veandco/go-sdl2/img"
import "github.com/veandco/go-sdl2/ttf"
import "github.com/veandco/go-sdl2/mix"
import "pjre/ScriptManager"
import "pjre/PJRDecompiler"
import "pjre/GameEngine"
/*import "pjre/Sections"
import "pjre/Informations"*/

func main() {
	// //flags
	gameFlags := GameEngine.NewGameFlags()

	if (len(os.Args) == 3) && (os.Args[2] == "-f") {
		gameFlags.ScreenMode = sdl.WINDOW_FULLSCREEN
	} else if len(os.Args) != 2 {
		fmt.Println("pjrengine <script path> [-f:fullscreen]\n")
		fmt.Println("Build 1. Created by Igor Santarek. MIT License.")
		os.Exit(1)
	}

	scriptMngr := ScriptManager.New()
	err := scriptMngr.LoadScript(os.Args[1])
	if err != nil {
		panic(err)
	}

	decompiler := PJRDecompiler.New(scriptMngr.GetByteCode())
	val, err := decompiler.DecompileSections()
	if err != nil {
		panic(err)
	}

	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	img.Init(img.INIT_PNG)
	defer img.Quit()

	if err := ttf.Init(); err != nil {
		panic(err)
	}
	defer ttf.Quit()

	if err := mix.OpenAudio(mix.DEFAULT_FREQUENCY, mix.DEFAULT_FORMAT, mix.DEFAULT_CHANNELS, 4096); err != nil {
		panic(err)
	}
	defer mix.CloseAudio()

	window, err := sdl.CreateWindow("Sample", 
		sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED,
		800, 600, gameFlags.ScreenMode)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	renderer, err := sdl.CreateRenderer(window, -1, 
		sdl.RENDERER_SOFTWARE)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()

	engine, err := GameEngine.New(val, window, renderer)
	if err != nil {
		panic(err)
	}
	defer engine.Destroy()

	err = engine.Loop()
	if err != nil {
		panic(err)
	}
}
