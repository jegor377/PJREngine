package ScriptManager

type ScriptManager struct {
	byteCode []byte
}

func New() *ScriptManager {
	return &ScriptManager{nil}
}

func (sm *ScriptManager) GetByteCode() []byte {
	return sm.byteCode
}