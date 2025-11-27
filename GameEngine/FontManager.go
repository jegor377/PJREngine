package GameEngine

import "github.com/veandco/go-sdl2/ttf"
import "github.com/veandco/go-sdl2/sdl"
import "errors"

type FontManager struct {
	fonts map[string]*ttf.Font
}

func NewFontManager() *FontManager {
	return &FontManager{make(map[string]*ttf.Font, 0)}
}

func (mgr *FontManager) Destroy() {
	for _, v := range mgr.fonts {
		v.Close()
	}
	mgr.fonts = nil
}

func (mgr *FontManager) LoadFont(name, fontPath string, size int) error {
	if _, ok := mgr.fonts[name]; ok {
		return errors.New("Cannot to load font. Font already exists: " + name)
	}
	font, err := ttf.OpenFont(fontPath, size)
	mgr.fonts[name] = font
	return err
}

func (mgr *FontManager) RemoveFont(name string) error {
	if v, ok := mgr.fonts[name]; ok {
		v.Close()
		mgr.fonts[name] = nil
		return nil
	}
	return errors.New("Cannot to remove font. Font not found.")
}

func (mgr *FontManager) GetTexture(renderer *sdl.Renderer, name, text string, color sdl.Color) (*sdl.Texture, error) {
	if text == "" {
		return nil, errors.New("Cannot get text. Text is empty.")
	}
	font, ok := mgr.fonts[name]
	if !ok {
		return nil, errors.New("Cannot get text. Font not found: " + name)
	}

	// Renderujemy tekst na surface
	surface, err := font.RenderUTF8Blended(text, color)
	if err != nil {
		return nil, err
	}
	defer surface.Free() // zwalniamy surface po zakończeniu

	// Tworzymy teksturę z surface
	texture, err := renderer.CreateTextureFromSurface(surface)
	if err != nil {
		return nil, err
	}

	return texture, nil
}
