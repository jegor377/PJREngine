package CommandMovingManager

type JumpingHandler interface {
	GetJump(byteCode []byte, iter uint32) uint32
}