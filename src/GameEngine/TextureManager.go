package GameEngine

import "github.com/veandco/go-sdl2/sdl"
import "github.com/veandco/go-sdl2/img"
import "errors"
import "strconv"

type TextureManager struct {
	textures map[uint32]*sdl.Texture
}

func NewTextureManager() *TextureManager {
	return &TextureManager{make(map[uint32]*sdl.Texture)}
}

func (tm *TextureManager) Destroy() {
	for _, v := range tm.textures {
		v.Destroy()
	}
	tm.textures = nil
}

func (tm *TextureManager) LoadTexture(renderer *sdl.Renderer, id uint32, path string) error {
	if path == "" {
		return errors.New("Cannot to load texture. Path is empty.")
	} else if path == "none" {
		tm.textures[0] = nil
		return nil
	}
	if _, ok := tm.textures[id]; !ok {
		texture, err := img.LoadTexture(renderer, path)
		tm.textures[id] = texture
		return err
	}
	return errors.New("Texture already exists: " + strconv.Itoa(int(id)))
}

func (tm *TextureManager) RemoveTexture(id uint32) error {
	if val, ok := tm.textures[id]; ok {
		val.Destroy()
		tm.textures[id] = nil
		delete(tm.textures, id)
		return nil
	}
	return errors.New("Cannot to remove texture. Texture id not found: " + strconv.Itoa(int(id)))
}

func (tm *TextureManager) GetTexture(id uint32) (*sdl.Texture, error) {
	if id == 0 {
		return nil, nil
	}
	if val, ok := tm.textures[id]; ok {
		return val, nil
	}
	return nil, errors.New( "Cannot to get texture. Texture id not found: " + strconv.Itoa(int(id)) )
}