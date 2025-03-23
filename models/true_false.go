package models

type TrueFalse struct {
	Value bool `json:"value" required:"true"`
}

// Vrátí nový objekt TrueFalse.
//
//	@param value
//	@return *TrueFalse
func NewTrueFalse(value bool) *TrueFalse {
	return &TrueFalse{Value: value}
}
