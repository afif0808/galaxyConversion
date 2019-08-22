package services

import "testing"

func TestGalaxyNumberService(t *testing.T) {
	//ConvertToArabicNumber TEST
	var galaxyNumberService = GalaxyNumberService{
		GalaxyNumberDatabase: map[string]*GalaxyNumber{
			"glob":  &GalaxyNumber{Name: "glob", Value: 1, Subtractable: true, Repeatable: true},
			"prok":  &GalaxyNumber{Name: "prok", Value: 5, Subtractable: false, Repeatable: false},
			"pish":  &GalaxyNumber{Name: "pish", Value: 10, Subtractable: true, Repeatable: true},
			"tegj":  &GalaxyNumber{Name: "tegj", Value: 50, Subtractable: false, Repeatable: false},
			"tus":   &GalaxyNumber{Name: "tus", Value: 100, Subtractable: true, Repeatable: true},
			"matus": &GalaxyNumber{Name: "matus", Value: 500, Subtractable: false, Repeatable: false},
			"bu":    &GalaxyNumber{Name: "bu", Value: 1000, Subtractable: true, Repeatable: true},
		},
		ArabicNumberDatabase: map[float64]string{
			1:    "glob",
			5:    "prok",
			10:   "pish",
			50:   "tegj",
			100:  "tus",
			500:  "matus",
			1000: "bu",
		},
		ArabicNumberLimit: 9800,
		GalaxyFractionNumberDatabase: map[string]float64{
			"dot": 1,
			"haf": 5,
		},
		ArabicFractionNumberDatabase: map[float64]string{
			1: "dot",
			5: "haf",
		},
	}
	number1, getNumber1Err := galaxyNumberService.ConvertToArabicNumber("glob")
	if number1 != 1 {
		t.Errorf("Error : expecting '2' got '%v'", number1)
	}
	if getNumber1Err != nil {
		t.Errorf("Error : expecting 'nil' error got '%v'", getNumber1Err)
	}

	number42, getNumber42Err := galaxyNumberService.ConvertToArabicNumber("pish-tegj-glob-glob")
	if number42 != 42 {
		t.Errorf("Error : expecting '42' got '%v'", number42)
	}
	if getNumber42Err != nil {
		t.Errorf("Error : expecting 'nil' error got '%v'", getNumber42Err)
	}

	// WITH FRACTIONS
	number1dot6, getNumber1dot6Err := galaxyNumberService.ConvertToArabicNumber("glob haf-dot")
	// t.Errorf("WHY %v", number)
	if number1dot6 != 1.6 {
		t.Errorf("Error : expecting '1.1' got '%v'", number1dot6)
	}
	if getNumber1dot6Err != nil {
		t.Errorf("Error : expecting 'nil' error got '%v'", getNumber1dot6Err)
	}

	//ERROR CONDITIONS

	//ERROR : UNKNOWN NUMBER
	number, unknownNumberErr := galaxyNumberService.ConvertToArabicNumber("one")
	if number != 0 {
		t.Errorf("Error : expecting '0' got '%v'", number)
	}
	if unknownNumberErr == nil {
		t.Errorf("Error : expecting 'error'  got 'nil' error")
	}

	//ERROR : EMPTY INPUT
	number, emptyInputErr := galaxyNumberService.ConvertToArabicNumber("")
	if number != 0 {
		t.Errorf("Error : expecting '0' got '%v'", number)
	}
	if emptyInputErr == nil {
		t.Errorf("Error : expecting 'error'  got 'nil' error")
	}

	//ERROR : NUMBER APPEAR MORE THAN 3 TIMES WITHOUT A SUBTRACTION BEFORE
	number, succession4TimesErr := galaxyNumberService.ConvertToArabicNumber("glob-glob-glob-glob")
	if number != 0 {
		t.Errorf("Error : expecting '0' got '%v'", number)
	}
	if succession4TimesErr == nil {
		t.Errorf("Error : expecting 'error'  got 'nil' error")
	}

	// ERRORS : EQUAL OR BIGGER NUMBER APPEAR AFTER SUBTRACTION

	number, equalNumberAfterSubtractionErr := galaxyNumberService.ConvertToArabicNumber("glob-prok-glob")
	if number != 0 {
		t.Errorf("Error : expecting '0' got '%v'", number)
	}
	if equalNumberAfterSubtractionErr == nil {
		t.Errorf("Error : expecting 'error'  got 'nil' error")
	}

	number, biggerNumberAfterSubtraction := galaxyNumberService.ConvertToArabicNumber("glob-prok-pish")
	if number != 0 {
		t.Errorf("Error : expecting '0' got '%v'", number)
	}
	if biggerNumberAfterSubtraction == nil {
		t.Errorf("Error : expecting 'error'  got 'nil' error")
	}

	//ERROR : UNREPEATABLE NUMBER - EXAMPLE : prok-prok
	number, unrepeatableErr := galaxyNumberService.ConvertToArabicNumber("prok-prok")
	if number != 0 {
		t.Errorf("Error : expecting '0' got '%v'", number)
	}
	if unrepeatableErr == nil {
		t.Errorf("Error : expecting 'error'  got 'nil' error")
	}

	//ERROR : UNSUBTRACTABLE NUMBER - EXAMPLE : prok-pish
	number, unsubtractableErr := galaxyNumberService.ConvertToArabicNumber("prok-pish")
	if number != 0 {
		t.Errorf("Error : expecting '0' got '%v'", number)
	}
	if unsubtractableErr == nil {
		t.Errorf("Error : expecting 'error'  got 'nil' error")
	}

	//ERROR : EXCEEDS THE SUBTRAHEND LIMIT
	number, exceedSubtrahendLimitErr := galaxyNumberService.ConvertToArabicNumber("glob-tegj")
	if number != 0 {
		t.Errorf("Error : expecting '0' got '%v'", number)
	}
	if exceedSubtrahendLimitErr == nil {
		t.Errorf("Error : expecting 'error'  got 'nil' error")
	}

	//ERROR : SUBTRACTION APPEAR AFTER EQUAL OR SMALLER NUMBER
	number, subtractionAppearAfterEqualOrSmallerNumberErr := galaxyNumberService.ConvertToArabicNumber("prok-glob-pish")
	if number != 0 {
		t.Errorf("Error : expecting '0' got '%v'", number)
	}
	if subtractionAppearAfterEqualOrSmallerNumberErr == nil {
		t.Errorf("Error : expecting 'error'  got 'nil' error")
	}

	//ConvertToGalaxyNumber TEST

	//ConvertToGalaxyNumber() VALID INPUTS

	//ConvertToGalaxyNumber() : GET GalaxyNumber of 1998
	GN1998, getGN1998Err := galaxyNumberService.ConvertToGalaxyNumber(1998)
	if GN1998 != "bu-tus-bu-pish-tus-prok-glob-glob-glob" {
		t.Errorf("Error : expected 'bu-tus-bu-pish-tus-prok-glob-glob-glob'  got  '%v'", GN1998)
	}
	if getGN1998Err != nil {
		t.Errorf("Error : expected 'error' got  'nil' error")
	}

	//ConvertToGalaxyNumber() WTIH FRACTIONS
	GN14dot1, getGN14dot1Err := galaxyNumberService.ConvertToGalaxyNumber(14.1)
	if GN14dot1 != "pish-glob-prok dot" {
		t.Errorf("Error : expected 'pish-glob-prok .'  got  '%v'", GN14dot1)
	}
	if getGN14dot1Err != nil {
		t.Errorf("Error : expected 'error' got  'nil' error")
	}
	//ConvertToGalaxyNumber() : GET GalaxyNumber of 1
	GN1, getGN1Err := galaxyNumberService.ConvertToGalaxyNumber(1)
	if GN1 != "glob" {
		t.Errorf("Error : expected 'pish'  got  '%v'", GN1)
	}
	if getGN1Err != nil {
		t.Errorf("Error : expected 'error' got  'nil' error")
	}
	//ConvertToGalaxyNumber() : GET GalaxyNumber of 4
	GN4, getGN4Err := galaxyNumberService.ConvertToGalaxyNumber(4)
	if GN4 != "glob-prok" {
		t.Errorf("Error : expected 'glob-prok'  got  '%v'", GN4)
	}
	if getGN4Err != nil {
		t.Errorf("Error : expected 'error' got  'nil' error")
	}
	//ConvertToGalaxyNumber() : GET GalaxyNumber of 5
	GN5, getGN5Err := galaxyNumberService.ConvertToGalaxyNumber(5)
	if GN5 != "prok" {
		t.Errorf("Error : expected 'prok'  got  '%v'", GN5)
	}
	if getGN5Err != nil {
		t.Errorf("Error : expected 'error' got  'nil' error")
	}
	//ConvertToGalaxyNumber() : GET GalaxyNumber of 8
	GN8, getGN8Err := galaxyNumberService.ConvertToGalaxyNumber(8)
	if GN8 != "prok-glob-glob-glob" {
		t.Errorf("Error : expected 'prok-glob-glob-glob'  got  '%v'", GN8)
	}
	if getGN8Err != nil {
		t.Errorf("Error : expected 'error' got  'nil' error")
	}
	//ConvertToGalaxyNumber() : GET GalaxyNumber of 9
	GN9, getGN9Err := galaxyNumberService.ConvertToGalaxyNumber(9)
	if GN9 != "glob-pish" {
		t.Errorf("Error : expected 'glob-pish'  got  '%v'", GN9)
	}
	if getGN9Err != nil {
		t.Errorf("Error : expected 'error' got  'nil' error")
	}
	//ConvertToGalaxyNumber() : GET GalaxyNumber of 10
	GN10, getGN10Err := galaxyNumberService.ConvertToGalaxyNumber(10)
	if GN10 != "pish" {
		t.Errorf("Error : expected 'pish'  got  '%v'", GN10)
	}
	if getGN10Err != nil {
		t.Errorf("Error : expected 'error' got  'nil' error")
	}

	//ConvertToGalaxyNumber() ERROR CONIDITIONS

	//ConvertToGalaxyNumber() ERROR : ZERO INPUT
	GN0, smallerThanOneInputErr := galaxyNumberService.ConvertToGalaxyNumber(0)
	if GN0 != "" {
		t.Errorf("Error : expected 'nil' got  %v", GN0)
	}
	if smallerThanOneInputErr == nil {
		t.Errorf("Error : expected 'error' got  'nil' error")
	}

	//ConvertToGalaxyNumber() ERROR : EXCEEDS THE NUMBER LIMIT - EXAMPLE : '11' assigned where the limit is 10
	GNBeyondLimit, exceedsNumberLimitErr := galaxyNumberService.ConvertToGalaxyNumber(
		galaxyNumberService.ArabicNumberLimit + 1,
	)
	if GNBeyondLimit != "" {
		t.Errorf("Error : expected 'nil' got  %v", GNBeyondLimit)
	}
	if exceedsNumberLimitErr == nil {
		t.Errorf("Error : expected 'error' got  'nil' error")
	}

	// ConvertToGalaxyNumberFractions TEST
	// ConvertToGalaxyNumberFractions VALID INPUTS
	// ConvertToGalaxyNumberFractions GET galaxyNumber of 0.4
	GFN0dot4, GFN0do4Err := galaxyNumberService.ConvertToGalaxyNumberFractions(0.4)
	if GFN0dot4 != "dot-dot-dot-dot" {
		t.Errorf("Error : expecting 'dot-dot-dot-dot' got %v", GFN0dot4)
	}
	if GFN0do4Err != nil {
		t.Errorf("Error : expecting 'nil' error got %v", GFN0do4Err)
	}

	// ConvertToGalaxyNumberFractions ERROR CONDITIONS
	// ConvertToGalaxyNumberFractions ERROR : ASSIGNING NON-FRAGMENT NUMBER
	GFN2, assigningNonFragmentNumberErr := galaxyNumberService.ConvertToGalaxyNumberFractions(2)
	if GFN2 != "" {
		t.Errorf("Error : expecting '' got '%v'", GFN2)
	}
	if assigningNonFragmentNumberErr == nil {
		t.Errorf("Error : expecting 'error' got 'nil' error")
	}

	// ConvertToArabicNumberFractions TEST
	// ConvertToArabicNumberFractions VALID INPUTS
	AFN0dot5, AFN0dot5Err := galaxyNumberService.ConvertToArabicFractionNumber("haf")
	if AFN0dot5 != 0.5 {
		t.Errorf("Error : expecting '0.5' got '%v'", AFN0dot5)
	}
	if AFN0dot5Err != nil {
		t.Errorf("Error : expecting 'nil' error got %v", AFN0dot5Err)
	}

	//ConvertToArabicFractionNumber ERROR : BIGGER NUMBER AFTER SMALLER NUMBER
	AFNdothaf, biggerNumberAfterSmallerNumberErr := galaxyNumberService.ConvertToArabicFractionNumber("dot-haf")
	if AFNdothaf != 0 {
		t.Errorf("Error : expecting '0' got '%v'", AFNdothaf)
	}
	if biggerNumberAfterSmallerNumberErr == nil {
		t.Errorf("Error : expecting  'error' got 'nil'")
	}

	// ConvertToArabicFractionNumber ERROR : INVALID GALAXY FRACTION NUMBER
	AFNhafdot5, inbalidFractionNumberErr := galaxyNumberService.ConvertToArabicFractionNumber("haf-dot-dot-dot-dot-dot")
	if AFNhafdot5 != 0 {
		t.Errorf("Error : expecting '0' got '%v'", inbalidFractionNumberErr)
	}
	if inbalidFractionNumberErr == nil {
		t.Errorf("Error : expecting  'error' got 'nil'")
	}

	// ConvertToArabicFractionNumber ERROR : EMPTY NUMBER
	AFNempty, AFNemptyError := galaxyNumberService.ConvertToArabicFractionNumber("")
	if AFNempty != 0 {
		t.Errorf("Error : expecting '0' got '%v'", AFNemptyError)
	}
	if AFNemptyError == nil {
		t.Errorf("Error : expecting  'error' got 'nil'")

	}

	// ConvertToArabicFractionNumber ERROR : NON GALAXY FRACTION NUMBER INPUT
	nonGalaxyFractionNumber, nonGalaxyFractionNumberErr := galaxyNumberService.ConvertToArabicFractionNumber("half-dotz")
	if nonGalaxyFractionNumber != 0 {
		t.Errorf("Error : expecting '0' got '%v'", nonGalaxyFractionNumber)
	}
	if nonGalaxyFractionNumberErr == nil {
		t.Errorf("Error : expecting  'error' got 'nil'")

	}

}
