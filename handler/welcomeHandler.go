package handler

import (
	"net/http"
)

type WelcomeHandler struct {}

func (handler *WelcomeHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	http.ServeFile(writer, request, "welcome.html")
}