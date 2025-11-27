package Sections

import "pjre/Informations"

type DialogOption struct {
	id uint32
	nextDialogId uint32
	byteCode []byte
	Info *Informations.BasicInformations
}

func NewDialogOption(id uint32, nextDialogId uint32, byteCode []byte) (*DialogOption, error) {
	dialogOption := DialogOption{id, nextDialogId, byteCode, Informations.NewBasic()}
	err := dialogOption.loadAttributes()
	return &dialogOption, err
}

func (option *DialogOption) GetId() uint32 {
	return option.id
}

func (option *DialogOption) GetTarget() uint32 {
	return option.nextDialogId
}
