package middleware

import (
	"fmt"
	"net/http"
	"runtime"

	"kelas-golang-pzn/go-restful-api/helper"
	"kelas-golang-pzn/go-restful-api/model/web"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("sampai kesini kah?")
	_, file, no, ok := runtime.Caller(1)
	if ok {
		fmt.Printf("called from %s#%d\n", file, no)
	}
	if "RAHASIA" == request.Header.Get("X-API-Key") {
		// ok
		fmt.Println("pass to next function")
		middleware.Handler.ServeHTTP(writer, request)
		return
	}

	// error
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusUnauthorized)

	webResponse := web.WebResponse{
		Code:   http.StatusUnauthorized,
		Status: "UNAUTHORIZED",
	}

	helper.WriteToResponseBody(writer, webResponse)
}
