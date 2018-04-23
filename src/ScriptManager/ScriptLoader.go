package ScriptManager

import "io/ioutil"
import "errors"

const CORRECT_PJR_SCRIPT_PREFIX = "PJR"

func (sm *ScriptManager) LoadScript(path string) error {
	byteCode, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	if !isPJRScript(byteCode[:3]) {
		return errors.New(path + " is not a PJR script.")
	}
	sm.byteCode = byteCode
	return nil
}

func isPJRScript(byteCode []byte) bool {
	return string(byteCode) == CORRECT_PJR_SCRIPT_PREFIX
}