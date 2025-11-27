package CommandFactory

import "pjre/Informations"

type SCommand interface {
	// returns new iterator, app state
	DoCmd(byteCode []byte, iter uint32, appState *Informations.AppState) (uint32, *Informations.AppState)
}
