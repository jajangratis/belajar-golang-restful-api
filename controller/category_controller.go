package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type CategoryController interface {
	Create(write http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(write http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(write http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindById(write http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAll(write http.ResponseWriter, request *http.Request, params httprouter.Params)
}
