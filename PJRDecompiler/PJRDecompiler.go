package PJRDecompiler

type PJRDecompiler struct {
	byteCode []byte
}

func New(byteCode []byte) *PJRDecompiler {
	return &PJRDecompiler{byteCode}
}
