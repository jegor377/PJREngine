package Commands

import "pjre/Informations"

type DTextCommand struct {
}

func NewDTextCommand() *DTextCommand {
	return &DTextCommand{}
}

func (cmd *DTextCommand) DoCmd(byteCode []byte, iter uint32, info *Informations.DialogInformations) (uint32, *Informations.DialogInformations) {
	var newText string = ""
	for iter < uint32( len(byteCode) ) {
		if byteCode[iter] == 0 {
			break
		}
		newText += string(byteCode[iter])
		iter++
	}
	info.Text = newText
	return iter + 1, info
}
