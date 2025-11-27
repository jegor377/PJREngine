package Commands

import "pjre/Informations"

type WYDTCommand struct {
}

func NewWYDTCommand() *WYDTCommand {
	return &WYDTCommand{}
}

func (cmd *WYDTCommand) DoCmd(byteCode []byte, iter uint32, appState *Informations.AppState) (uint32, *Informations.AppState) {
	var newWYDT string = ""
	for iter < uint32( len(byteCode) ) {
		if byteCode[iter] == 0 {
			break
		}
		newWYDT += string(byteCode[iter])
		iter++
	}
	appState.WhatYouDoText.Change(newWYDT)
	return iter + 1, appState
}
