package Commands

import "Informations"
import "ConvertingTools"

type OptionCommand struct {
}

func NewOptionCommand() *OptionCommand {
	return &OptionCommand{}
}

// TODO: Popraw to koniecznie! Errory mają wychodzić wszystkie do maina! :P
func (cmd *OptionCommand) DoCmd(byteCode []byte, iter uint32, info *Informations.DialogInformations) (uint32, *Informations.DialogInformations) {
	var dialogOptionsId []byte
	iter, dialogOptionsId = getDialogOptionId(byteCode, iter, iter + 4)
	convDialogOptionsId, err := ConvertingTools.ConvertBytesToUint32LittleEndian(dialogOptionsId)
	if err != nil {
		panic(err)
	}
	info.AddOption(convDialogOptionsId)
	return iter, info
}