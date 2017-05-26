package server

import (
	"server/handler"
	"log"
	"github.com/gorilla/mux"
)

var (
	wHandler 	= &handler.WelcomeHandler{}
	jsHandler	= &handler.JavaScriptHandler{}
	phpHandler	= &handler.PhpHandler{}
	fsHandler	= &handler.FileSystemHandler{}

	router		*mux.Router
)

func RegisterRoutes(mux *mux.Router) {
	router = mux

	folderPathPattern := `/{_dummy:[^\?\%\*\:\|\"<>]*}`
	jsFilePattern := `/{_dummy:[^\?\%\*\:\|\"<>]+\.js}`
	phpFilePattern := `/{_dummy:[^\?\%\*\:\|\"<>]+\.php}`

	createRoute("/welcome", wHandler)
	createRoute(jsFilePattern, jsHandler)
	createRoute(phpFilePattern, phpHandler)
	createRoute(folderPathPattern, fsHandler)
}

func createRoute(pattern string, handler handler.IHandler) {
	if router == nil {
		log.Fatal("Access DENIED!")
	}

	router.NewRoute().Path(pattern).Handler(handler)

}