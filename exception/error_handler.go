package exception

import (
	"belajar-golang-resful-api/helper"
	"belajar-golang-resful-api/model/web"
	"net/http"

	"github.com/go-playground/validator"
)

func ErrorHandler(write http.ResponseWriter, request *http.Request, err any) {

	if notFoundError(write, request, err) {
		return
	}

	if validationErrors(write, request, err) {
		return
	}

	internalServerError(write, request, err)
}

func validationErrors(write http.ResponseWriter, request *http.Request, err any) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		write.Header().Set("Content-Type", "application/json")
		write.WriteHeader(http.StatusBadRequest)
	
		webResponse := web.WebResponse{
			Code: http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data: exception.Error(),
		}

		helper.WriteToResponseBody(write, webResponse)
		return true
	} else {
		return false
	}

}

func notFoundError(write http.ResponseWriter, request *http.Request, err any) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		write.Header().Set("Content-Type", "application/json")
		write.WriteHeader(http.StatusNotFound)
	
		webResponse := web.WebResponse{
			Code: http.StatusNotFound,
			Status: "NOT FOUND",
			Data: exception.Error,
		}

		helper.WriteToResponseBody(write, webResponse)
		return true
	} else {
		return false
	}
}

func internalServerError(write http.ResponseWriter, request *http.Request, err any) {
	write.Header().Set("Content-Type", "application/json")
	write.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code: http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data: err,
	
	}

	helper.WriteToResponseBody(write, webResponse)
}
