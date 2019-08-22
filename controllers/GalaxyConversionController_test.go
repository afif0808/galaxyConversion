package controllers

import (
	"encoding/json"
	"galaxyConversion/exchangeObjects"
	"galaxyConversion/interfaces"
	"galaxyConversion/services"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestConvert(t *testing.T) {
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

	var galaxyConversionController = &GalaxyConversionController{
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

	router := mux.NewRouter()
	router.HandleFunc("/convert/{convertFrom}/{convertTo}/{amount}", galaxyConversionController.Convert)

	successRequest := httptest.NewRequest("GET", "http://localhost:343/convert/credit/gold/20004", nil)
	succcessButReturnEmptyRequest := httptest.NewRequest("GET", "http://localhost:343/convert/silver/gold/glob", nil)
	unknownObjectRequest := httptest.NewRequest("GET", "http://localhost:343/convert/emerald/coral/glob", nil)
	badRequest := httptest.NewRequest("GET", "http://localhost:343/convert/credit/gold/blabablabl", nil)
	w := httptest.NewRecorder()

	var jsonResponse map[string]interface{}

	// ERROR : SUCCESS REQUEST
	router.ServeHTTP(w, successRequest)
	json.NewDecoder(w.Body).Decode(&jsonResponse)
	if status := jsonResponse["status"]; status != "success" {
		t.Errorf("Error : expecting 'success' status got '%v'", status)
	}

	// ERROR : SUCCESS BUT RETURN EMPTY REQUEST
	router.ServeHTTP(w, succcessButReturnEmptyRequest)
	json.NewDecoder(w.Body).Decode(&jsonResponse)
	if statusCode := jsonResponse["statusCode"]; statusCode != float64(204) {
		t.Errorf("Error : expecting '204' status got '%v'", statusCode)
	}

	// ERROR : UNKNOWN OBJECT REQUEST
	router.ServeHTTP(w, unknownObjectRequest)
	json.NewDecoder(w.Body).Decode(&jsonResponse)
	if status := jsonResponse["status"]; status != "error" {
		t.Errorf("Error : expecting 'error' status got '%v'", status)
	}

	// ERROR : ANOTHER BAD REQUEST
	router.ServeHTTP(w, badRequest)
	json.NewDecoder(w.Body).Decode(&jsonResponse)
	if status := jsonResponse["status"]; status != "error" {
		t.Errorf("Error : expecting 'error' status got '%v'", status)
	}

}
