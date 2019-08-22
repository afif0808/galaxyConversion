package exchangeObjects

import (
	"galaxyConversion/interfaces"
)

type Gold struct {
	Value float64
	interfaces.IGalaxyNumberService
}

func (gold *Gold) GetObjectValue() float64 {
	return gold.Value
}
func (gold *Gold) GetFloat64Amount(amount string) (float64, error) {
	var float64Amount float64

	float64Amount, convertToArabicNumberErr := gold.ConvertToArabicNumber(amount)
	if convertToArabicNumberErr != nil {
		return 0, convertToArabicNumberErr
	}
	float64Amount *= gold.Value
	return float64Amount, nil
}
func (gold *Gold) GetAmountInObjectFormat(amount float64) (interface{}, error) {
	return gold.ConvertToGalaxyNumber(amount)

}
