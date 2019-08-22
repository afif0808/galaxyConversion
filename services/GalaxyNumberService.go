package services

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type GalaxyNumberService struct {
	GalaxyNumberDatabase         map[string]*GalaxyNumber
	GalaxyFractionNumberDatabase map[string]float64
	ArabicNumberDatabase         map[float64]string
	ArabicFractionNumberDatabase map[float64]string
	ArabicNumberLimit            float64
}

func (gn *GalaxyNumberService) ConvertToGalaxyNumber(arabicNumber float64) (string, error) {
	if arabicNumber <= 0 {
		return "", fmt.Errorf("Error : cannot assign 0 or smaller")
	}
	if arabicNumber > gn.ArabicNumberLimit {
		return "", fmt.Errorf("Error : %v exceeds the number limit : %v", arabicNumber, gn.ArabicNumberLimit)
	}

	var fractions float64
	var galaxyFractionNumbers string
	var galaxyNumber string
	var arabicNumberStr string

	//Separating number and its fraction
	arabicNumber, fractions = math.Modf(arabicNumber)

	//Converting arabic fraction number to galaxy fraction number
	//If fraction is bigger than '0' and smaller than '0'
	if fractions >= 0.05 && fractions < 1 {
		galaxyFractionNumbers, _ = gn.ConvertToGalaxyNumberFractions(fractions)
		// if arabic number is '0' return the fraction only
	}

	if arabicNumber == 0 {
		return galaxyFractionNumbers, nil
	}

	//Converting arabicNumber(float64) to string
	arabicNumberStr = fmt.Sprint(arabicNumber)
	//Iterate through 'arabicNumberStr' to treat each non-zero digit separately
	// EXAMPLE : '1998' to '1000' , '900' , '90' , '8'
	for i, numberInRune := range arabicNumberStr {
		//'digitCount' is the count of digit after the number
		//'pow10' is 10 to the power of 'digitCount'
		var numberPosition = i + 1
		var digitCount int = len(arabicNumberStr) - numberPosition
		var pow10 float64 = math.Pow10(digitCount)
		var baseNumber, _ = strconv.ParseFloat(string(numberInRune), 64)
		var number = baseNumber * pow10
		var one float64 = 1 * pow10
		var three float64 = 3 * pow10
		var five float64 = 5 * pow10
		var eight float64 = 8 * pow10
		var ten float64 = 10 * pow10
		//
		switch true {
		case number < five && number <= three:
			for i := 0; i < int(baseNumber); i++ {
				galaxyNumber += gn.ArabicNumberDatabase[one] + "-"
			}
		case number < five && number > three:
			galaxyNumber += gn.ArabicNumberDatabase[one] + "-" + gn.ArabicNumberDatabase[five] + "-"
		case number == five:
			galaxyNumber += gn.ArabicNumberDatabase[five] + "-"
		case number > five && number <= eight:
			galaxyNumber += gn.ArabicNumberDatabase[five] + "-"
			for i := 0; i < 3; i++ {
				galaxyNumber += gn.ArabicNumberDatabase[one] + "-"
			}
		case number > five && number > eight:
			galaxyNumber += gn.ArabicNumberDatabase[one] + "-" +
				gn.ArabicNumberDatabase[ten] + "-"
		}
	}

	galaxyNumber = galaxyNumber[:len(galaxyNumber)-1]
	if fractions >= 0.05 && fractions < 1 {
		galaxyNumber += " " + galaxyFractionNumbers
	}
	return galaxyNumber, nil
}

func (gn *GalaxyNumberService) ConvertToArabicNumber(galaxyNumbersStr string) (float64, error) {
	var galaxyNumbersStrSplit = strings.Split(galaxyNumbersStr, " ")
	var galaxyFractionNumber string
	var arabicNumber float64
	var galaxyNumbers = &GalaxyNumberNode{}
	var galaxyNumberAppearances = map[string]int{}
	var galaxyNumbersArray = strings.Split(galaxyNumbersStrSplit[0], "-")

	if len(galaxyNumbersArray) == 0 {
		return 0, fmt.Errorf("Error : empty input ")
	}

	if len(galaxyNumbersStrSplit) > 1 {
		galaxyFractionNumber = galaxyNumbersStrSplit[1]
		arabicFractionNumber, convertToArabicFractionNumberErr := gn.ConvertToArabicFractionNumber(galaxyFractionNumber)
		if convertToArabicFractionNumberErr != nil {
			return 0, convertToArabicFractionNumberErr
		}
		arabicNumber += arabicFractionNumber
	}
	for i, v := range galaxyNumbersArray {
		var galaxyNumber *GalaxyNumber
		var found bool
		galaxyNumber, found = gn.GalaxyNumberDatabase[v]
		if found == false {
			return 0, fmt.Errorf("Error : There is no number '%v' ", v)
		}
		if i == 0 {
			galaxyNumbers = &GalaxyNumberNode{
				galaxyNumber: galaxyNumber,
				value:        galaxyNumber.Value,
			}

			galaxyNumberAppearances[galaxyNumber.Name] += 1
			galaxyNumbers.tail = galaxyNumbers
		} else {
			addNextNumberErr := galaxyNumbers.AddNextNode(
				&GalaxyNumberNode{
					galaxyNumber: galaxyNumber,
					value:        galaxyNumber.Value,
				},
				galaxyNumberAppearances[galaxyNumber.Name],
			)
			galaxyNumberAppearances[galaxyNumber.Name] += 1

			if addNextNumberErr != nil {
				return 0, addNextNumberErr
			}
		}
	}

	arabicNumber += galaxyNumbers.GetArabicNumber()
	return arabicNumber, nil
}



func (gn *GalaxyNumberService) ConvertToGalaxyNumberFractions(fractions float64) (string, error) {
	if fractions >= 1 {
		return "", fmt.Errorf("Error : %v is not a fraction number", fractions)
	}
	fractions = math.Round(fractions*100) / 100
	fractions, _ = strconv.ParseFloat(string(fmt.Sprint(fractions)[2]), 64)
	var galaxyFractionNumberStr string
	switch true {
	case fractions > 5:
		galaxyFractionNumberStr += gn.ArabicFractionNumberDatabase[5] + "-"
		for i := 0; i < int(fractions)-5; i++ {
			galaxyFractionNumberStr += gn.ArabicFractionNumberDatabase[1] + "-"
		}
	case fractions == 5:
		galaxyFractionNumberStr += gn.ArabicFractionNumberDatabase[5] + "-"
	case fractions < 5:
		for i := 0; i < int(fractions); i++ {
			galaxyFractionNumberStr += gn.ArabicFractionNumberDatabase[1] + "-"
		}
	}
	return galaxyFractionNumberStr[:len(galaxyFractionNumberStr)-1], nil
}
func (gn *GalaxyNumberService) ConvertToArabicFractionNumber(galaxyFractionNumber string) (float64, error) {
	if galaxyFractionNumber == "" {
		return 0, fmt.Errorf("Error : empty ")
	}
	var galaxyFractionNumbersArray = strings.Split(galaxyFractionNumber, "-")
	var arabicFractionNumber float64
	for i, number := range galaxyFractionNumbersArray {
		var numberNext string
		var arabicNumber float64
		var arabicNumberNext float64

		arabicNumber, found := gn.GalaxyFractionNumberDatabase[number]
		if found == false {
			return 0, fmt.Errorf("Error : '%v' is  not a galaxy fraction number", arabicNumber)
		}
		if len(galaxyFractionNumbersArray)-1 != i {
			numberNext = galaxyFractionNumbersArray[i+1]
			arabicNumberNext, found = gn.GalaxyFractionNumberDatabase[numberNext]
			if found == false {
				return 0, fmt.Errorf("Error : '%v' is  not a galaxy fraction number", arabicNumberNext)
			}
			if arabicNumberNext > arabicNumber {
				return 0, fmt.Errorf("Error : bigger number appear after smaller number")
			}
		}
		arabicFractionNumber += arabicNumber
		// After sum all number
		if arabicFractionNumber > 9 {
			return 0, fmt.Errorf("Error : '%v' is not valid galaxy fraction number", galaxyFractionNumber)
		}

	}
	arabicFractionNumber /= 10
	return arabicFractionNumber, nil
}
