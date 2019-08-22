package exchangeObjects

import (
	"galaxyConversion/interfaces"
)

type Silver struct {
	Value float64
	interfaces.IGalaxyNumberService
}

func (silver *Silver) GetObjectValue() float64 {
	return silver.Value
}
func (silver *Silver) GetFloat64Amount(amount string) (float64, error) {
	var float64Amount float64

	float64Amount, convertToArabicNumberErr := silver.ConvertToArabicNumber(amount)
	if convertToArabicNumberErr != nil {
		return 0, convertToArabicNumberErr
	}
	float64Amount *= silver.Value
	return float64Amount, nil
}
func (silver *Silver) GetAmountInObjectFormat(amount float64) (interface{}, error) {
	return silver.ConvertToGalaxyNumber(amount)
}
