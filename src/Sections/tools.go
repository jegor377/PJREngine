package Sections

func isCommandAllowed(cmdOpcode uint8, allowedCommandOpcodes []uint8) bool {
	for _, v := range allowedCommandOpcodes {
		if cmdOpcode == v {
			return true
		}
	}
	return false
}

func getCommandOpcode(byteCode []byte, iter uint32) uint8 {
	return uint8(byteCode[iter])
}