package interfaces

type IGalaxyConversionService interface {
	Convert(convertFrom, convertTo IObject, amount string) (interface{}, error)
}
