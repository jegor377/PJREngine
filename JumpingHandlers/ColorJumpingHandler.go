package JumpingHandlers

type ColorJumpingHandler struct {
}

func NewColorJumpingHandler() *ColorJumpingHandler {
	return &ColorJumpingHandler{}
}

func (jmpH *ColorJumpingHandler) GetJump(byteCode []byte, iter uint32) uint32 {
	return iter + 1
}
