package models

type TrueFalse struct {
	Value bool `json:"value"`
}

// Vrátí nový objekt TrueFalse.
//
//	@param value
//	@return *TrueFalse
func NewTrueFalse(value bool) *TrueFalse {
	return &TrueFalse{Value: value}
}
