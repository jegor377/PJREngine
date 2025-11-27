package Sections

import "pjre/Informations"
import "pjre/CommandFactory"
import "pjre/CommandMovingManager"

// text, color
func (option *DialogOption) loadAttributes() error {
	allowedCommandOpcodes := []uint8{
		3,
		4,
	}
	var iter uint32 = 0 // TODO: nie będzie działać, bo nie przeskakuje komend, których nie zna
	for iter < uint32(len(option.byteCode)) { // basicly do the command if it is in allowed. If not, then jump to the next command if it is not the end of byte code.
		cmdOpcode := getCommandOpcode(option.byteCode, iter)
		iter += 1
		if isCommandAllowed(cmdOpcode, allowedCommandOpcodes) {
			cmdFactory := CommandFactory.NewDO()
			cmd, err := cmdFactory.Get(cmdOpcode)
			if err != nil {
				return err
			}
			iter, option.Info = cmd.DoCmd(option.byteCode, iter, option.Info)
		} else {
			cmdMovingManager := CommandMovingManager.New()
			cmd, err := cmdMovingManager.Get(cmdOpcode)
			if err != nil {
				return err
			}
			iter = cmd.GetJump(option.byteCode, iter)
		}
	}
	return nil
}

// bgColor, appName, wydt, bgSong
func (option *DialogOption) DoJob(appState *Informations.AppState) (*Informations.AppState, error) {
	allowedCommandOpcodes := []uint8{
		0,
		1,
		2,
		8,
	}
	var iter uint32 = 0
	var _appState *Informations.AppState = appState
	for iter < uint32(len(option.byteCode)) { // basicly do the command if it is in allowed. If not, then jump to the next command if it is not the end of byte code.
		cmdOpcode := getCommandOpcode(option.byteCode, iter)
		iter += 1
		if isCommandAllowed(cmdOpcode, allowedCommandOpcodes) {
			cmdFactory := CommandFactory.NewS()
			cmd, err := cmdFactory.Get(cmdOpcode)
			if err != nil {
				return nil, err
			}
			iter, _appState = cmd.DoCmd(option.byteCode, iter, _appState)
		} else {
			cmdMovingManager := CommandMovingManager.New()
			cmd, err := cmdMovingManager.Get(cmdOpcode)
			if err != nil {
				return nil, err
			}
			iter = cmd.GetJump(option.byteCode, iter)
		}
	}
	_appState.SetActualDialogId(option.nextDialogId)
	return _appState, nil
}
