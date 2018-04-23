package JumpingHandlers

type BgColorJumpingHandler struct {
}

func NewBgColorJumpingHandler() *BgColorJumpingHandler {
	return &BgColorJumpingHandler{}
}

func (jmpH *BgColorJumpingHandler) GetJump(byteCode []byte, iter uint32) uint32 {
	return iter + 1
}