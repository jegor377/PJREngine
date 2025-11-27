package GameEngine

import "fmt"

// songs
func (engine *GameEngine) loadSounds() {
	if !engine.songs.IsEmpty {
		size := float32( len(engine.songs.Songs) )
		var actualState float32 = 0
		fmt.Println("Loading sounds...")
		for id, v := range engine.songs.Songs {
			if err := engine.chunkMgr.LoadChunk(uint32(id), v); err != nil {
				panic(err)
			}
			actualState++
			fmt.Print("Loading sounds [", int( (actualState/size) * 100.0 ), "%]\n")
		}
	}
}

func (engine *GameEngine) playSound(soundId uint32) {
	if err := engine.chunkMgr.PlayChunk(soundId); err != nil {
		panic(err)
	}
}
