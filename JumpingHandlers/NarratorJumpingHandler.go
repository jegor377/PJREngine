package JumpingHandlers

type NarratorJumpingHandler struct {
}

func NewNarratorJumpingHandler() *NarratorJumpingHandler {
	return &NarratorJumpingHandler{}
}

func (jmpH *NarratorJumpingHandler) GetJump(byteCode []byte, iter uint32) uint32 {
	return iter + 4
}
