package Commands

import "pjre/Informations"

type DOColorCommand struct {
}

func NewDOColorCommand() *DOColorCommand {
	return &DOColorCommand{}
}

func (cmd *DOColorCommand) DoCmd(byteCode []byte, iter uint32, info *Informations.BasicInformations) (uint32, *Informations.BasicInformations) {
	var newColor uint8 = 0
	if iter < uint32( len(byteCode) ) {
		newColor = uint8(byteCode[iter])
		iter++
	}
	info.Color = newColor
	return iter, info
}
