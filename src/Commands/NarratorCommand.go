package Commands

import "Informations"
import "ConvertingTools"

type NarratorCommand struct {
}

func NewNarratorCommand() *NarratorCommand {
	return &NarratorCommand{}
}

// TODO: Popraw to koniecznie! Errory mają wychodzić wszystkie do maina! :P
func (cmd *NarratorCommand) DoCmd(byteCode []byte, iter uint32, info *Informations.DialogInformations) (uint32, *Informations.DialogInformations) {
	var songId []byte
	iter, songId = getSongId(byteCode, iter, iter + 4)
	convSongId, err := ConvertingTools.ConvertBytesToUint32LittleEndian(songId)
	if err != nil {
		panic(err)
	}
	info.Narrator = convSongId + 1
	return iter, info
}