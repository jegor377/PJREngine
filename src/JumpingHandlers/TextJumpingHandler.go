package JumpingHandlers

type TextJumpingHandler struct {
}

func NewTextJumpingHandler() *TextJumpingHandler {
	return &TextJumpingHandler{}
}

func (jmpH *TextJumpingHandler) GetJump(byteCode []byte, iter uint32) uint32 {
	for iter < uint32( len(byteCode) ) {
		if byteCode[iter] == 0 {
			break
		}
		iter++
	}
	return iter + 1
}