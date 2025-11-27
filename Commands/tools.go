package Commands

func getSongId(byteCode []byte, iter uint32, idSizeOffset uint32) (uint32, []byte) {
	var songId []byte = make([]byte, 0)
	for ( iter < uint32( len(byteCode) ) ) && ( iter < idSizeOffset ) {
		songId = append(songId, byteCode[iter])
		iter++
	}
	return iter, songId
}

func getDialogOptionId(byteCode []byte, iter uint32, idSizeOffset uint32) (uint32, []byte) {
	var dialogOptionId []byte = make([]byte, 0)
	for ( iter < uint32( len(byteCode) ) ) && ( iter < idSizeOffset ) {
		dialogOptionId = append(dialogOptionId, byteCode[iter])
		iter++
	}
	return iter, dialogOptionId
}
