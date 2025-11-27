package Commands

import "pjre/Informations"

type DOTextCommand struct {
}

func NewDOTextCommand() *DOTextCommand {
	return &DOTextCommand{}
}

func (cmd *DOTextCommand) DoCmd(byteCode []byte, iter uint32, info *Informations.BasicInformations) (uint32, *Informations.BasicInformations) {
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
