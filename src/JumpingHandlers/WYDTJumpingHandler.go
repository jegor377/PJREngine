package JumpingHandlers

type WYDTJumpingHandler struct {
}

func NewWYDTJumpingHandler() *WYDTJumpingHandler {
	return &WYDTJumpingHandler{}
}

func (jmpH *WYDTJumpingHandler) GetJump(byteCode []byte, iter uint32) uint32 {
	for iter < uint32( len(byteCode) ) {
		if byteCode[iter] == 0 {
			break
		}
		iter++
	}
	return iter + 1
}