package JumpingHandlers

type AppNameJumpingHandler struct {
}

func NewAppNameJumpingHandler() *AppNameJumpingHandler {
	return &AppNameJumpingHandler{}
}

func (jmpH *AppNameJumpingHandler) GetJump(byteCode []byte, iter uint32) uint32 {
	for iter < uint32( len(byteCode) ) {
		if byteCode[iter] == 0 {
			break
		}
		iter++
	}
	return iter + 1
}
