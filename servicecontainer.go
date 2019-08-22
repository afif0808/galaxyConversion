package main

import (
	"galaxyConversion/controllers"
	"galaxyConversion/exchangeObjects"
	"galaxyConversion/interfaces"
	"galaxyConversion/services"
	"sync"
)

type IServiceContainer interface {
	InjectGalaxyConversionController() controllers.GalaxyConversionController
}

type kernel struct{}

func (k *kernel) InjectGalaxyConversionController() controllers.GalaxyConversionController {
	var galaxyNumberService = services.GalaxyNumberService{
		GalaxyNumberDatabase: map[string]*services.GalaxyNumber{
			"glob":  &services.GalaxyNumber{Name: "glob", Value: 1, Subtractable: true, Repeatable: true},
			"prok":  &services.GalaxyNumber{Name: "prok", Value: 5, Subtractable: false, Repeatable: false},
			"pish":  &services.GalaxyNumber{Name: "pish", Value: 10, Subtractable: true, Repeatable: true},
			"tegj":  &services.GalaxyNumber{Name: "tegj", Value: 50, Subtractable: false, Repeatable: false},
			"tus":   &services.GalaxyNumber{Name: "tus", Value: 100, Subtractable: true, Repeatable: true},
			"matus": &services.GalaxyNumber{Name: "matus", Value: 500, Subtractable: false, Repeatable: false},
			"bu":    &services.GalaxyNumber{Name: "bu", Value: 1000, Subtractable: true, Repeatable: true},
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

	var galaxyConversionController = controllers.GalaxyConversionController{
		Objects: map[string]interfaces.IObject{
			"gold": &exchangeObjects.Gold{
				Value:                14450,
				IGalaxyNumberService: &galaxyNumberService,
			},
			"iron": &exchangeObjects.Iron{
				Value:                195.5,
				IGalaxyNumberService: &galaxyNumberService,
			},
			"silver": &exchangeObjects.Silver{
				Value:                17,
				IGalaxyNumberService: &galaxyNumberService,
			},
			"credit": &exchangeObjects.Credit{
				Value: 1,
			},
		},
		IGalaxyConversionService: &services.ConversionService{},
	}

	return galaxyConversionController
}

var (
	k             *kernel
	containerOnce sync.Once
)

func ServiceContainer() IServiceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
		})
	}
	return k
}
