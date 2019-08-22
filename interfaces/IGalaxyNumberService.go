package interfaces

type IGalaxyNumberService interface {
	ConvertToArabicNumber(galaxyNumbersInput string) (float64, error)
	ConvertToGalaxyNumber(arabicNumber float64) (string, error)
}
