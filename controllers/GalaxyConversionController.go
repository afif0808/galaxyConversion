package controllers

import (
	"encoding/json"
	"fmt"
	"galaxyConversion/interfaces"
	"net/http"

	"github.com/gorilla/mux"
)

type GalaxyConversionController struct {
	interfaces.IGalaxyConversionService
	Objects map[string]interfaces.IObject
}

func (gcc *GalaxyConversionController) Convert(w http.ResponseWriter, r *http.Request) {
	var objects map[string]interfaces.IObject = gcc.Objects

	var convertFrom string = mux.Vars(r)["convertFrom"]
	var convertTo string = mux.Vars(r)["convertTo"]
	var amount string = mux.Vars(r)["amount"]
	var result interface{}
	var jsonResponse map[string]interface{}
	var jsonEncoder = json.NewEncoder(w)
	var convertFromObject interfaces.IObject
	var convertToObject interfaces.IObject
	var notFoundObjects string

	w.Header().Set("Content-Type", "application/json")

	convertFromObject, convertFromObjectFound := objects[convertFrom]
	if convertFromObjectFound == false {
		notFoundObjects += fmt.Sprintf("'%v' ", convertFrom)
	}

	convertToObject, convertToObjectFound := objects[convertTo]
	if convertFromObjectFound == false {
		notFoundObjects += fmt.Sprintf("'%v' ", convertFrom)
	}

	if convertFromObjectFound == false || convertToObjectFound == false {
		w.WriteHeader(400)
		jsonResponse = map[string]interface{}{
			"status":     "error",
			"statusCode": 400,
			"message":    fmt.Sprintf("%v not found", notFoundObjects),
			"data":       nil,
		}
		jsonEncoder.Encode(jsonResponse)
		return
	}

	result, convertErr := gcc.IGalaxyConversionService.Convert(convertFromObject, convertToObject, amount)
	if convertErr != nil {
		w.WriteHeader(400)
		jsonResponse = map[string]interface{}{
			"status":     "error",
			"statusCode": 400,
			"message":    convertErr.Error(),
			"data":       nil,
		}
		jsonEncoder.Encode(jsonResponse)
		return
	}

	if result == "" || result == float64(0) {
		w.WriteHeader(204)
		jsonResponse = map[string]interface{}{
			"status":     "success",
			"statusCode": 204,
			"message":    "",
			"data":       result,
		}
		jsonEncoder.Encode(jsonResponse)
		return
	}

	jsonResponse = map[string]interface{}{
		"status":     "success",
		"statusCode": 200,
		"message":    "",
		"data":       result,
	}
	w.WriteHeader(200)
	jsonEncoder.Encode(jsonResponse)
}
