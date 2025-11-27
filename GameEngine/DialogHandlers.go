package GameEngine

import "errors"
import "strconv"
import "pjre/Sections"

func (engine *GameEngine) findDialog(dialogId uint32) (*Sections.Dialog, error) {
	var dialog *Sections.Dialog = nil
	for _, v := range engine.dialogs.Dialogs {
		if v.GetId() == dialogId {
			dialog = v
		}
	}
	if dialog == nil {
		return nil, errors.New("Dialog not found: " + strconv.Itoa(int(dialogId)))
	}
	return dialog, nil
}

func (engine *GameEngine) loadActualDialog() error {
	dialogId := engine.appState.GetActualDialogId()
	dialog, err := engine.findDialog(dialogId)
	if err != nil {
		return err
	}
	if err := engine.changeActualSprite(dialog.Info.Sprite); err != nil {
		return err
	}
	if err := engine.changeActualText(dialog.Info); err != nil {
		return err
	}
	engine.playSound(dialog.Info.Narrator)
	engine.buttonMgr.Clear()
	if !dialog.IsEnding() {
		engine.loadOptions()
	}
	engine.textMouseOffset = 0
	return nil
}

func (engine *GameEngine) checkForActualDialog() error {
	if engine.appState.HasDialogBeenChanged() {
		return engine.loadActualDialog()
	}
	return nil
}
