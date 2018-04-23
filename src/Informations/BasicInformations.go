package Informations

type BasicInformations struct {
	Color uint8
	Text string
}

func NewBasic() *BasicInformations {
	return &BasicInformations{0, ""}
}