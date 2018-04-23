package Sections

import "Informations"

type Dialog struct {
	id uint32
	ending bool
	byteCode []byte
	Info *Informations.DialogInformations
}

func NewDialog(id uint32, ending bool, byteCode []byte) (*Dialog, error) {
	dialog := Dialog{id, ending, byteCode, Informations.NewDialogInfo()}
	err := dialog.loadAttributes()
	return &dialog, err
}

func (dialog *Dialog) GetId() uint32 {
	return dialog.id
}

func (dialog *Dialog) IsEnding() bool {
	return dialog.ending
}