package Sections

import "errors"
import "pjre/ConvertingTools"

type DialogsSection struct {
	Dialogs []*Dialog
}

func NewDialogsSection() *DialogsSection {
	return &DialogsSection{nil}
}

func (dialogs *DialogsSection) DecodeDialogs(byteCode []byte) error {
	if len(byteCode) == 0 {
		return errors.New("Cannot to decode empty dialogs bytecode.")
	}
	var iter uint32 = 0
	for iter < uint32(len(byteCode)) {
		var id uint32
		var ending bool
		var err error
		var code []byte

		id, err = getDialogOptionId(byteCode, iter)
		if err != nil {
			return err
		}
		iter += 4
		ending, err = getDialogEndingFlag(byteCode, iter)
		if err != nil {
			return err
		}
		iter += 1
		code, iter, err = getDialogCode(byteCode, iter)
		if err != nil {
			return err
		}
		dialog, err := NewDialog(id, ending, code)
		if err != nil {
			return err
		}
		dialogs.Dialogs = append(dialogs.Dialogs, dialog)
	}
	return nil
}

func getDialogId(byteCode []byte, iter uint32) (value uint32, err error) {
	id := byteCode[iter:iter + 4]
	value, err = ConvertingTools.ConvertBytesToUint32LittleEndian(id)
	return
}

func getDialogEndingFlag(byteCode []byte, iter uint32) (bool, error) {
	if byteCode[iter] > 0 {
		return true, nil
	}
	return false, nil
}

func getDialogSize(byteCode []byte, iter uint32) (value uint32, err error) {
	s := byteCode[iter:iter + 4]
	value, err = ConvertingTools.ConvertBytesToUint32LittleEndian(s)
	return value, err
}

func getDialogCode(byteCode []byte, iter uint32) (code []byte, newIterator uint32, err error) {
	size, err := getDialogSize(byteCode, iter)
	iter += 4
	newIterator = iter + size
	code = byteCode[iter:newIterator]
	return
}
