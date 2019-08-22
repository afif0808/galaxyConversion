package exchangeObjects

import (
	"fmt"
	"strconv"
)

type Credit struct {
	Value float64
}

func (credit *Credit) GetObjectValue() float64 {
	return credit.Value
}
func (credit *Credit) GetFloat64Amount(amount string) (float64, error) {
	amountInt, parseIntErr := strconv.ParseFloat(amount, 64)
	if parseIntErr != nil {
		return 0, fmt.Errorf("Error : expecting 'arabic number' got '%v'", amount)
	}
	return amountInt, nil
}
func (credit *Credit) GetAmountInObjectFormat(amount float64) (interface{}, error) {
	return amount * 1, nil
}
