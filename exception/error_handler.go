package exception

import (
	"github.com/go-playground/validator"
	"jajangratis/belajar-golang-restful-api/helper"
	"jajangratis/belajar-golang-restful-api/model/web"
	"net/http"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, error interface{}) {
	if notFoundError(writer, request, error) {
		return
	}
	if validationErrors(writer, request, error) {
		return
	}
	internalServerError(writer, request, error)
}

func validationErrors(writer http.ResponseWriter, request *http.Request, error interface{}) bool {
	exception, ok := error.(validator.ValidationErrors)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   exception.Error(),
		}

		helper.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

func notFoundError(writer http.ResponseWriter, request *http.Request, error interface{}) bool {
	exception, ok := error.(NotFoundError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Data Not Found",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, request *http.Request, error interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "Internal Server Error",
		Data:   error,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
