package Commands

import "pjre/Informations"
import "pjre/ConvertingTools"

type BgSongCommand struct {
}

func NewBgSongCommand() *BgSongCommand {
	return &BgSongCommand{}
}

// TODO: Popraw to koniecznie! Errory mają wychodzić wszystkie do maina! :P
func (cmd *BgSongCommand) DoCmd(byteCode []byte, iter uint32, appState *Informations.AppState) (uint32, *Informations.AppState) {
	var songId []byte
	iter, songId = getSongId(byteCode, iter, iter + 4)
	convSongId, err := ConvertingTools.ConvertBytesToUint32LittleEndian(songId)
	if err != nil {
		panic(err)
	}
	appState.BgSong.Change(convSongId + 1)
	return iter, appState
}
