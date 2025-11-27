package JumpingHandlers

type BgSongJumpingHandler struct {
}

func NewBgSongJumpingHandler() *BgSongJumpingHandler {
	return &BgSongJumpingHandler{}
}

func (jmpH *BgSongJumpingHandler) GetJump(byteCode []byte, iter uint32) uint32 {
	return iter + 4
}
