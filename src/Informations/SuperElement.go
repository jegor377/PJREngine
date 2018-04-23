package Informations

type Element interface{}

type SuperElement struct {
	element Element
	hasChanged bool
}

func NewSuperElement(element Element) *SuperElement {
	return &SuperElement{element, false}
}

func (se *SuperElement) HasChanged() bool {
	return se.hasChanged
}

func (se *SuperElement) Change(element Element) {
	se.element = element
	se.hasChanged = true
}

func (se *SuperElement) Get() Element {
	se.hasChanged = false
	return se.element
}