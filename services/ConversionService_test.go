package services

import (
	"galaxyConversion/exchangeObjects"
	"testing"
)

func TestConversionService(t *testing.T) {
	var galaxyNumberConversionService = GalaxyNumberService{
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
	var gold = exchangeObjects.Gold{
		Value:                14450,
		IGalaxyNumberService: &galaxyNumberConversionService,
	}
	var credit = exchangeObjects.Credit{
		Value: 1,
	}
	var silver = exchangeObjects.Silver{
		Value:                17,
		IGalaxyNumberService: &galaxyNumberConversionService,
	}
	var iron = exchangeObjects.Iron{
		Value:                195.5,
		IGalaxyNumberService: &galaxyNumberConversionService,
	}
	// var convertFrom = credit
	// var convertTo = gold
	var conversionService = ConversionService{}
	creditToGold, creditToGoldErr := conversionService.Convert(&credit, &gold, "14450")
	if creditToGold != "glob" {
		t.Errorf("Error : expecting 'glob' got '%v' ", creditToGold)
	}
	if creditToGoldErr != nil {
		t.Errorf("Error : expecting no error got %v", creditToGoldErr)
	}

	goldToCredit, goldToCreditErr := conversionService.Convert(&gold, &credit, "glob-prok")
	if goldToCredit != float64(57800) {
		t.Errorf("Error : expecting '57800' got '%v' ", goldToCredit)
	}
	if goldToCreditErr != nil {
		t.Errorf("Error : expecting no error got %v", goldToCreditErr)
	}

	ironToCredit, ironToCreditErr := conversionService.Convert(&iron, &credit, "glob-prok")
	if ironToCredit != float64(782) {
		t.Errorf("Error : expecting '782' got '%v' ", ironToCredit)
	}
	if ironToCreditErr != nil {
		t.Errorf("Error : expecting no error got %v", ironToCreditErr)
	}

	creditToIron, creditToIronErr := conversionService.Convert(&credit, &iron, "586.5")
	if creditToIron != "glob-glob-glob" {
		t.Errorf("Error : expecting 'glob-glob-glob' got '%v' ", creditToIron)
	}
	if creditToIronErr != nil {
		t.Errorf("Error : expecting no error got %v", creditToIronErr)
	}

	silverToCredit, silverToCreditErr := conversionService.Convert(&silver, &credit, "glob-prok")
	if silverToCredit != float64(68) {
		t.Errorf("Error : expecting '68' got '%v' ", silverToCredit)
	}
	if silverToCreditErr != nil {
		t.Errorf("Error : expecting no error got %v", silverToCreditErr)
	}

	creditToSilver, creditToSilverErr := conversionService.Convert(&credit, &silver, "50")
	if creditToSilver != "glob-glob haf-dot-dot-dot-dot" {
		t.Errorf("Error : expecting 'glob-glob haf-dot-dot-dot-dot' got '%v' ", creditToSilver)
	}
	if creditToSilverErr != nil {
		t.Errorf("Error : expecting no error got %v", creditToSilverErr)
	}

	goldToSilver, goldToSilverErr := conversionService.Convert(&gold, &silver, "glob")
	if goldToSilver != "matus-tus-tus-tus-tegj" {
		t.Errorf("Error : expecting 'matus-tus-tus-tus-tegj' got '%v' ", goldToSilver)
	}
	if goldToSilverErr != nil {
		t.Errorf("Error : expecting no error got %v", goldToSilverErr)
	}

	silverToGold, silverToGoldErr := conversionService.Convert(&silver, &gold, "bu")
	if silverToGold != "glob dot" {
		t.Errorf("Error : expecting 'glob-dot' got '%v' ", silverToGold)
	}
	if silverToGoldErr != nil {
		t.Errorf("Error : expecting no error got %v", silverToGoldErr)
	}

}
