package CommandFactory

import "pjre/Informations"

type DCommand interface {
	// returns new iterator, basic informations
	DoCmd(byteCode []byte, iter uint32, info *Informations.DialogInformations) (uint32, *Informations.DialogInformations)
}
