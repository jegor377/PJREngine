package Sections

import "Informations"
import "CommandFactory"
import "CommandMovingManager"

type SettingsSection struct {
}

func NewSettingsSection() *SettingsSection {
	return &SettingsSection{}
}

func (settings *SettingsSection) DoJob(byteCode []byte, appState *Informations.AppState) (*Informations.AppState, error) {
	allowedCommandOpcodes := []uint8{
		0,
		1,
		2,
		8,
	}
	var iter uint32 = 0
	var _appState *Informations.AppState = appState
	for iter < uint32(len(byteCode)) { // basicly do the command if it is in allowed. If not, then jump to the next command if it is not the end of byte code.
		cmdOpcode := getCommandOpcode(byteCode, iter)
		iter += 1
		if isCommandAllowed(cmdOpcode, allowedCommandOpcodes) {
			cmdFactory := CommandFactory.NewS()
			cmd, err := cmdFactory.Get(cmdOpcode)
			if err != nil {
				return nil, err
			}
			iter, _appState = cmd.DoCmd(byteCode, iter, _appState)
		} else {
			cmdMovingManager := CommandMovingManager.New()
			cmd, err := cmdMovingManager.Get(cmdOpcode)
			if err != nil {
				return nil, err
			}
			iter = cmd.GetJump(byteCode, iter)
		}
	}
	return _appState, nil
}