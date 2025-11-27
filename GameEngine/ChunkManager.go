package GameEngine

import "github.com/veandco/go-sdl2/mix"
import "strconv"
import "errors"

type ChunkManager struct {
	chunks map[uint32]*mix.Chunk
	actualBgChannel int
}

func NewChunkManager() *ChunkManager {
	return &ChunkManager{
		make(map[uint32]*mix.Chunk, 0),
		-1,
	}
}

func (chunkMgr *ChunkManager) Destroy() {
	for _, v := range chunkMgr.chunks {
		v.Free()
	}
}

func (chunkMgr *ChunkManager) LoadChunk(id uint32, path string) (err error) {
	if id == 0 {
		chunkMgr.chunks[0] = nil
		return nil
	}
	if _, ok := chunkMgr.chunks[id]; !ok {
		chunkMgr.chunks[id], err = mix.LoadWAV(path)
		return
	}
	return errors.New("Chunk already exists: " + strconv.Itoa(int(id)))
}

func (chunkMgr *ChunkManager) RemoveChunk(id uint32) error {
	if v, ok := chunkMgr.chunks[id]; ok {
		v.Free()
		chunkMgr.chunks[id] = nil
		delete(chunkMgr.chunks, id)
		return nil
	}
	return errors.New("Could not find chunk: " + strconv.Itoa(int(id)) + ". Chunk does not exist.")
}

func (chunkMgr *ChunkManager) PlayChunk(chunkId uint32) error {
	if chunkId == 0 {
		return nil
	}
	if val, ok := chunkMgr.chunks[chunkId]; ok {
		val.Play(-1, 0)
		return nil
	}
	return errors.New("Could not play chunk because it does not exist. Game is broken.")
}

func (chunkMgr *ChunkManager) PlayRepeated(chunkId uint32, repeat int) (err error) {
	if chunkId == 0 {
		return nil
	}
	mix.HaltChannel(chunkMgr.actualBgChannel)
	if val, ok := chunkMgr.chunks[chunkId]; ok {
		chunkMgr.actualBgChannel, err = val.PlayTimed(-1, repeat, -1)
		mix.Volume(chunkMgr.actualBgChannel, mix.MAX_VOLUME / 4)
		return err
	}
	return errors.New("Could not play chunk because it does not exist. Game is broken.")
}
