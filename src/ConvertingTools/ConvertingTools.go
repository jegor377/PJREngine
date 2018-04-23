package ConvertingTools

import "bytes"
import "encoding/binary"

func ConvertUint32ToBytesLittleEndian(val uint32) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, val)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func ConvertBytesToUint32LittleEndian(val []byte) (uint32, error) {
	var ret uint32
	buf := bytes.NewReader(val)
	err := binary.Read(buf, binary.LittleEndian, &ret)
	if err != nil {
		return 0, err
	}
	return ret, nil
}