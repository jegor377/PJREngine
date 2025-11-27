package JumpingHandlers

type SpriteJumpingHandler struct {
}

func NewSpriteJumpingHandler() *SpriteJumpingHandler {
	return &SpriteJumpingHandler{}
}

func (jmpH *SpriteJumpingHandler) GetJump(byteCode []byte, iter uint32) uint32 {
	return iter + 4
}
