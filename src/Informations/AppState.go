package Informations

type AppState struct {
	WhatYouDoText *SuperElement //string, tak wiem, że źle napisane :p
	AppName *SuperElement // string
	BgColor *SuperElement // uint8
	BgSong *SuperElement // uint32
	IsEnding bool // bool
	actualDialog *SuperElement //uint32
}

func New() *AppState {
	return &AppState{
		NewSuperElement("What do you do?"), // wydt
		NewSuperElement("Sample"), // appName
		NewSuperElement(0xf), // bgColor
		NewSuperElement(0), // bgSong
		false,
		NewSuperElement(0), // dialogId
	}
}

func (appState *AppState) HasDialogBeenChanged() bool {
	return appState.actualDialog.HasChanged()
}

func (appState *AppState) SetActualDialogId(newDialog uint32) {
	appState.actualDialog.Change(newDialog)
}

func (appState *AppState) GetActualDialogId() uint32 {
	val, _ := appState.actualDialog.Get().(uint32)
	return val
}