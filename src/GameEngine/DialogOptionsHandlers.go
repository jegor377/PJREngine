package GameEngine

import "errors"
import "strconv"
import "Sections"

func (engine *GameEngine) findDialogOption(dialogOptionId uint32) (*Sections.DialogOption, error) {
	var dialogOption *Sections.DialogOption = nil
	for _, v := range engine.dialogOptions.DialogOptions {
		if v.GetId() == dialogOptionId {
			dialogOption = v
		}
	}
	if dialogOption == nil {
		return nil, errors.New("DialogOption not found: " + strconv.Itoa(int(dialogOptionId)))
	}
	return dialogOption, nil
}

func (engine *GameEngine) loadOptions() {
	dialogId := engine.appState.GetActualDialogId()
	dialog, err := engine.findDialog(dialogId)
	if err != nil {
		panic(err)
	}
	for _, v := range dialog.Info.Options {
		dialogOption, err := engine.findDialogOption(v)
		if err != nil {
			panic(err)
		}
		engine.buttonMgr.Add(dialogOption.Info.Text, dialogOption.GetTarget(), dialogOption.GetId(), dialogOption.Info.Color, engine.fontMgr)
	}
}