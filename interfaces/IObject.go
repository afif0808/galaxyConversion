package interfaces

type IObject interface {
	GetObjectValue() float64
	GetFloat64Amount(amount string) (float64, error)
	GetAmountInObjectFormat(amount float64) (interface{}, error)
}
