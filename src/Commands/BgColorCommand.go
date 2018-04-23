package Commands

import "Informations"

type BgColorCommand struct {
}

func NewBgColorCommand() *BgColorCommand {
	return &BgColorCommand{}
}

func (cmd *BgColorCommand) DoCmd(byteCode []byte, iter uint32, appState *Informations.AppState) (uint32, *Informations.AppState) {
	var newBgColor uint8 = 0
	if iter < uint32( len(byteCode) ) {
		newBgColor = uint8(byteCode[iter])
		iter++
	}
	appState.BgColor.Change(newBgColor)
	return iter, appState
}