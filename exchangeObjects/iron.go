package exchangeObjects

import (
	"galaxyConversion/interfaces"
)

type Iron struct {
	Value float64
	interfaces.IGalaxyNumberService
}

func (iron *Iron) GetObjectValue() float64 {
	return iron.Value
}
func (iron *Iron) GetFloat64Amount(amount string) (float64, error) {
	var float64Amount float64

	float64Amount, convertToArabicNumberErr := iron.ConvertToArabicNumber(amount)
	if convertToArabicNumberErr != nil {
		return 0, convertToArabicNumberErr
	}
	float64Amount *= iron.Value
	return float64Amount, nil
}
func (iron *Iron) GetAmountInObjectFormat(amount float64) (interface{}, error) {
	return iron.ConvertToGalaxyNumber(amount)
}
