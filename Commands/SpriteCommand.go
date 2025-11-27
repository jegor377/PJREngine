package Commands

import "pjre/Informations"
import "pjre/ConvertingTools"

type SpriteCommand struct {
}

func NewSpriteCommand() *SpriteCommand {
	return &SpriteCommand{}
}

func (cmd *SpriteCommand) DoCmd(byteCode []byte, iter uint32, info *Informations.DialogInformations) (uint32, *Informations.DialogInformations) {
	var spriteId []byte
	iter, spriteId = getSongId(byteCode, iter, iter + 4)
	convSpriteId, err := ConvertingTools.ConvertBytesToUint32LittleEndian(spriteId)
	if err != nil {
		panic(err)
	}
	info.Sprite = convSpriteId + 1
	return iter, info
}
