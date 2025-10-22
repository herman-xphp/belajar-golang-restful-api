package middleware

import (
	"belajar-golang-resful-api/helper"
	"belajar-golang-resful-api/model/web"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware  {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(write http.ResponseWriter, request *http.Request) {
	if "RAHASIA" == request.Header.Get("X-API-Key") {
		// ok
		middleware.Handler.ServeHTTP(write, request)
	} else {
		// error
		write.Header().Set("Content-Type", "application/json")
		write.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code: http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.WriteToResponseBody(write, webResponse)
	}
}
