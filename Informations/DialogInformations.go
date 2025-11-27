package Informations

type DialogInformations struct {
	*BasicInformations
	Options []uint32
	Narrator uint32
	Sprite uint32
}

func NewDialogInfo() *DialogInformations {
	return &DialogInformations{NewBasic(), nil, 0, 0}
}

func (dialogInfo *DialogInformations) AddOption(optionId uint32) {
	if !dialogInfo.hasOption(optionId) {
		if len(dialogInfo.Options) < 5 {
			dialogInfo.Options = append(dialogInfo.Options, optionId)
		}
	}
}

func (dialogInfo *DialogInformations) hasOption(optionId uint32) bool {
	for _, v := range dialogInfo.Options {
		if v == optionId {
			return true
		}
	}
	return false
}
