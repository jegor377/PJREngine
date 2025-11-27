package Commands

import "pjre/Informations"

type AppNameCommand struct {
}

func NewAppNameCommand() *AppNameCommand {
	return &AppNameCommand{}
}

func (cmd *AppNameCommand) DoCmd(byteCode []byte, iter uint32, appState *Informations.AppState) (uint32, *Informations.AppState) {
	var newName string = ""
	for iter < uint32( len(byteCode) ) {
		if byteCode[iter] == 0 {
			break
		}
		newName += string(byteCode[iter])
		iter++
	}
	appState.AppName.Change(newName)
	return iter + 1, appState
}
