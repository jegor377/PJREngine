package CommandFactory

import "Informations"

type DOCommand interface {
	// returns new iterator, basic informations
	DoCmd(byteCode []byte, iter uint32, info *Informations.BasicInformations) (uint32, *Informations.BasicInformations)
}