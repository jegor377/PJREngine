package PJRDecompiler

import "errors"
import "pjre/ConvertingTools"

func (decompiler *PJRDecompiler) DecompileSections() (map[string][]byte, error) {
	if len(decompiler.byteCode) < 23 {
		return nil, errors.New("Byte code is not valid.")
	}
	sections := make(map[string][]byte, 0)
	sprites, err := decompiler.GetSpritesSectionByteCode()
	if err != nil {
		return nil, err
	}
	songs, err := decompiler.GetSongsSectionByteCode()
	if err != nil {
		return nil, err
	}
	settings, err := decompiler.GetSettingsSectionByteCode()
	if err != nil {
		return nil, err
	}
	dialogOptions, err := decompiler.GetDialogOptionsSectionByteCode()
	if err != nil {
		return nil, err
	}
	dialogs, err := decompiler.GetDialogsSectionByteCode()
	sections["sprites"] = sprites
	sections["songs"] = songs
	sections["settings"] = settings
	sections["dialogOptions"] = dialogOptions
	sections["dialogs"] = dialogs
	return sections, err
}

func (decompiler *PJRDecompiler) GetSpritesSize() (uint32, error) {
	return ConvertingTools.ConvertBytesToUint32LittleEndian(decompiler.byteCode[3:7])
}

func (decompiler *PJRDecompiler) GetSongsSize() (uint32, error) {
	return ConvertingTools.ConvertBytesToUint32LittleEndian(decompiler.byteCode[7:11])
}

func (decompiler *PJRDecompiler) GetSettingsSize() (uint32, error) {
	return ConvertingTools.ConvertBytesToUint32LittleEndian(decompiler.byteCode[11:15])
}

func (decompiler *PJRDecompiler) GetDialogOptionsSize() (uint32, error) {
	return ConvertingTools.ConvertBytesToUint32LittleEndian(decompiler.byteCode[15:19])
}

func (decompiler *PJRDecompiler) GetDialogsSize() (uint32, error) {
	return ConvertingTools.ConvertBytesToUint32LittleEndian(decompiler.byteCode[19:23])
}

func (decompiler *PJRDecompiler) GetByteCode() ([]byte, error) {
	bSize := len(decompiler.byteCode)
	if bSize > 0 {
		return decompiler.byteCode[23:bSize - 1], nil
	}
	return nil, errors.New("Byte code is empty.")
}

func (decompiler *PJRDecompiler) getSectionByteCode(offset uint32, size uint32) ([]byte, error) {
	bSize := len(decompiler.byteCode)
	if bSize > 0 {
		return decompiler.byteCode[offset:size + offset], nil
	}
	return nil, errors.New("Byte code is empty.")
}

func (decompiler *PJRDecompiler) GetSpritesSectionByteCode() ([]byte, error) {
	spritesSize, err := decompiler.GetSpritesSize()
	if err != nil {
		return nil, err
	} 
	return decompiler.getSectionByteCode(23, spritesSize)
}

func (decompiler *PJRDecompiler) GetSongsSectionByteCode() ([]byte, error) {
	spritesSize, err := decompiler.GetSpritesSize()
	if err != nil {
		return nil, err
	}
	songsSize, err := decompiler.GetSongsSize()
	if err != nil {
		return nil, err
	}
	return decompiler.getSectionByteCode(23 + spritesSize, songsSize)
}

func (decompiler *PJRDecompiler) GetSettingsSectionByteCode() ([]byte, error) {
	spritesSize, err := decompiler.GetSpritesSize()
	if err != nil {
		return nil, err
	}
	songsSize, err := decompiler.GetSongsSize()
	if err != nil {
		return nil, err
	}
	settingsSize, err := decompiler.GetSettingsSize()
	if err != nil {
		return nil, err
	}
	offset := 23 + spritesSize + songsSize
	return decompiler.getSectionByteCode(offset, settingsSize)
}

func (decompiler *PJRDecompiler) GetDialogOptionsSectionByteCode() ([]byte, error) {
	spritesSize, err := decompiler.GetSpritesSize()
	if err != nil {
		return nil, err
	}
	songsSize, err := decompiler.GetSongsSize()
	if err != nil {
		return nil, err
	}
	settingsSize, err := decompiler.GetSettingsSize()
	if err != nil {
		return nil, err
	}
	dialogOptionsSize, err := decompiler.GetDialogOptionsSize()
	if err != nil {
		return nil, err
	}
	offset := 23 + spritesSize + songsSize + settingsSize
	return decompiler.getSectionByteCode(offset, dialogOptionsSize)
}

func (decompiler *PJRDecompiler) GetDialogsSectionByteCode() ([]byte, error) {
	spritesSize, err := decompiler.GetSpritesSize()
	if err != nil {
		return nil, err
	}
	songsSize, err := decompiler.GetSongsSize()
	if err != nil {
		return nil, err
	}
	settingsSize, err := decompiler.GetSettingsSize()
	if err != nil {
		return nil, err
	}
	dialogOptionsSize, err := decompiler.GetDialogOptionsSize()
	if err != nil {
		return nil, err
	}
	dialogsSize, err := decompiler.GetDialogsSize()
	if err != nil {
		return nil, err
	}
	offset := 23 + spritesSize + songsSize + settingsSize + dialogOptionsSize
	return decompiler.getSectionByteCode(offset, dialogsSize)
}
