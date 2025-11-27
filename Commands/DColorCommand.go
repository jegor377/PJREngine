package Commands

import "pjre/Informations"

type DColorCommand struct {
}

func NewDColorCommand() *DColorCommand {
	return &DColorCommand{}
}

func (cmd *DColorCommand) DoCmd(byteCode []byte, iter uint32, info *Informations.DialogInformations) (uint32, *Informations.DialogInformations) {
	var newColor uint8 = 0
	if iter < uint32( len(byteCode) ) {
		newColor = uint8(byteCode[iter])
		iter++
	}
	info.Color = newColor
	return iter, info
}
