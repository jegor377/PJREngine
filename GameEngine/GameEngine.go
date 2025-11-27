package GameEngine

import "pjre/Sections"
import "pjre/Informations"
import "github.com/veandco/go-sdl2/sdl"

const MAX_TEXT_SIZE int32 = 230

type GameEngine struct {
	sprites *Sections.SpritesSection
	songs *Sections.SongsSection
	settings *Sections.SettingsSection
	dialogOptions *Sections.DialogOptionsSection
	dialogs *Sections.DialogsSection

	appState *Informations.AppState
	textureMgr *TextureManager
	fontMgr *FontManager
	buttonMgr *ButtonManager
	chunkMgr *ChunkManager

	window *sdl.Window
	renderer *sdl.Renderer

	actualSprite *sdl.Texture
	actualText *sdl.Texture
	btnSprite *sdl.Texture

	quit bool
	textMouseOffset int32
}

func New(sectionsCode map[string][]byte, window *sdl.Window, renderer *sdl.Renderer) (*GameEngine, error) {
	engine := GameEngine{
		Sections.NewSpritesSection(),
		Sections.NewSongsSection(),
		Sections.NewSettingsSection(),
		Sections.NewDialogOptionsSection(),
		Sections.NewDialogsSection(),
		Informations.New(),
		NewTextureManager(),
		NewFontManager(),
		NewButtonManager(0),
		NewChunkManager(),
		window,
		renderer,
		nil,
		nil,
		nil,
		false,
		0,
	}
	engine.sprites.DecodeSprites(sectionsCode["sprites"])
	err := engine.loadTextures()
	if err != nil {
		return nil, err
	}

	engine.songs.DecodeSongs(sectionsCode["songs"])
	engine.loadSounds()

	engine.appState, err = engine.settings.DoJob(sectionsCode["settings"], engine.appState)
	if err != nil {
		return nil, err
	}
	
	if err := engine.dialogOptions.DecodeDialogOptions(sectionsCode["dialogOptions"]); err != nil {
		return nil, err
	}
	if err := engine.dialogs.DecodeDialogs(sectionsCode["dialogs"]); err != nil {
		return nil, err
	}

	if err := engine.fontMgr.LoadFont("default", "assets/default.ttf", 16); err != nil {
		panic(err)
	}
	if err := engine.fontMgr.LoadFont("option", "assets/default.ttf", 18); err != nil {
		panic(err)
	}
	engine.loadBtnSprite()

	return &engine, engine.loadActualDialog()
}

func (engine *GameEngine) Destroy() {
	engine.sprites = nil
	engine.songs = nil
	engine.dialogOptions = nil
	engine.dialogs = nil
	engine.textureMgr.Destroy()
	engine.fontMgr.Destroy()
	engine.buttonMgr.Destroy()
	engine.chunkMgr.Destroy()
}
