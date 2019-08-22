package services

import (
	"galaxyConversion/interfaces"
)

type ConversionService struct {
	convertFrom interfaces.IObject
	convertTo   interfaces.IObject
}

func (service *ConversionService) Convert(convertFrom, convertTo interfaces.IObject, amount string) (interface{}, error) {
	var float64Amount float64
	var getFloat64AmountErr error
	var convertToObjectValue float64
	var result float64

	service.convertFrom = convertFrom
	service.convertTo = convertTo

	float64Amount, getFloat64AmountErr = service.convertFrom.GetFloat64Amount(amount)
	convertToObjectValue = service.convertTo.GetObjectValue()
	result = float64Amount / convertToObjectValue
	if getFloat64AmountErr != nil {
		return nil, getFloat64AmountErr
	}
	resultInObjectFormat, getAmountInObjectFormatErr := service.convertTo.GetAmountInObjectFormat(result)
	if getAmountInObjectFormatErr != nil {
		return 0, getAmountInObjectFormatErr
	}
	return resultInObjectFormat, nil
}
