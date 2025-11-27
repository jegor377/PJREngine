package Sections

import "pjre/Informations"
import "pjre/CommandFactory"
import "pjre/CommandMovingManager"

func (dialog *Dialog) loadAttributes() error {
	allowedCommandOpcodes := []uint8{
		3,
		4,
		5,
		6,
		7,
	}
	var iter uint32 = 0 // TODO: nie będzie działać, bo nie przeskakuje komend, których nie zna
	for iter < uint32(len(dialog.byteCode)) { // basicly do the command if it is in allowed. If not, then jump to the next command if it is not the end of byte code.
		cmdOpcode := getCommandOpcode(dialog.byteCode, iter)
		iter += 1
		if isCommandAllowed(cmdOpcode, allowedCommandOpcodes) {
			cmdFactory := CommandFactory.NewD()
			cmd, err := cmdFactory.Get(cmdOpcode)
			if err != nil {
				return err
			}
			iter, dialog.Info = cmd.DoCmd(dialog.byteCode, iter, dialog.Info)
		} else {
			cmdMovingManager := CommandMovingManager.New()
			cmd, err := cmdMovingManager.Get(cmdOpcode)
			if err != nil {
				return err
			}
			iter = cmd.GetJump(dialog.byteCode, iter)
		}
	}
	return nil
}

func (dialog *Dialog) DoJob(appState *Informations.AppState) (*Informations.AppState, error) {
	allowedCommandOpcodes := []uint8{
		0,
		1,
		2,
		8,
	}
	var iter uint32 = 0
	var _appState *Informations.AppState = appState
	for iter < uint32(len(dialog.byteCode)) { // basicly do the command if it is in allowed. If not, then jump to the next command if it is not the end of byte code.
		cmdOpcode := getCommandOpcode(dialog.byteCode, iter)
		iter += 1
		if isCommandAllowed(cmdOpcode, allowedCommandOpcodes) {
			cmdFactory := CommandFactory.NewS()
			cmd, err := cmdFactory.Get(cmdOpcode)
			if err != nil {
				return nil, err
			}
			iter, _appState = cmd.DoCmd(dialog.byteCode, iter, _appState)
		} else {
			cmdMovingManager := CommandMovingManager.New()
			cmd, err := cmdMovingManager.Get(cmdOpcode)
			if err != nil {
				return nil, err
			}
			iter = cmd.GetJump(dialog.byteCode, iter)
		}
	}
	return _appState, nil
}
