package Sections

import "pjre/ConvertingTools"

type DialogOptionsSection struct {
	DialogOptions []*DialogOption
	IsEmpty bool
}

func NewDialogOptionsSection() *DialogOptionsSection {
	return &DialogOptionsSection{nil, false}
}

func (dialogOptions *DialogOptionsSection) DecodeDialogOptions(byteCode []byte) error {
	if len(byteCode) == 0 {
		dialogOptions.IsEmpty = true
		return nil
	}
	var iter uint32 = 0
	for iter < uint32(len(byteCode)) {
		var id uint32
		var nextId uint32
		var err error
		var code []byte

		id, err = getDialogOptionId(byteCode, iter)
		if err != nil {
			return err
		}
		iter += 4
		nextId, err = getDialogOptionNextDialog(byteCode, iter)
		if err != nil {
			return err
		}
		iter += 4
		code, iter, err = getDialogOptionCode(byteCode, iter)
		if err != nil {
			return err
		}
		dialogOption, err := NewDialogOption(id, nextId, code)
		if err != nil {
			return err
		}
		dialogOptions.DialogOptions = append(dialogOptions.DialogOptions, dialogOption)
	}
	return nil
}

func getDialogOptionId(byteCode []byte, iter uint32) (value uint32, err error) {
	id := byteCode[iter:iter + 4]
	value, err = ConvertingTools.ConvertBytesToUint32LittleEndian(id)
	return
}

func getDialogOptionNextDialog(byteCode []byte, iter uint32) (value uint32, err error) {
	id := byteCode[iter:iter + 4]
	value, err = ConvertingTools.ConvertBytesToUint32LittleEndian(id)
	return
}

func getDialogOptionSize(byteCode []byte, iter uint32) (value uint32, err error) {
	s := byteCode[iter:iter + 4]
	value, err = ConvertingTools.ConvertBytesToUint32LittleEndian(s)
	return value, err
}

func getDialogOptionCode(byteCode []byte, iter uint32) (code []byte, newIterator uint32, err error) {
	size, err := getDialogOptionSize(byteCode, iter)
	iter += 4
	newIterator = iter + size
	code = byteCode[iter:newIterator]
	return
}
