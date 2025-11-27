package JumpingHandlers

type OptionJumpingHandler struct {
}

func NewOptionJumpingHandler() *OptionJumpingHandler {
	return &OptionJumpingHandler{}
}

func (jmpH *OptionJumpingHandler) GetJump(byteCode []byte, iter uint32) uint32 {
	return iter + 4
}
