package main

import (
	"sync"

	"github.com/gorilla/mux"
)

type IMyRouter interface {
	InitRouter() *mux.Router
}

type router struct{}

func (router *router) InitRouter() *mux.Router {

	galaxyConversionController := ServiceContainer().InjectGalaxyConversionController()

	r := mux.NewRouter()
	r.HandleFunc("/convert/{convertFrom}/{convertTo}/{amount}", galaxyConversionController.Convert)

	return r
}

var (
	m          *router
	routerOnce sync.Once
)

func MyRouter() IMyRouter {
	if m == nil {
		routerOnce.Do(func() {
			m = &router{}
		})
	}
	return m
}
