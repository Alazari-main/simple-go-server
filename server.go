package server

import (
	"net/http"
	"server/config"
	"strconv"
	"github.com/gorilla/mux"
)

type Server struct{
		AppBuilder  IAppBuilder
		Rtr *mux.Router
}

func (server *Server) Start() {
	server.AppBuilder = NewAppBuilder()
	server.Rtr = mux.NewRouter()

	RegisterMiddlewares(server.AppBuilder)
	RegisterRoutes(server.Rtr)

	http.ListenAndServe(":"+strconv.Itoa(config.Instance.Port), &serverHandler{AppBuilder: server.AppBuilder, Rtr: server.Rtr})
}

type serverHandler struct{
		AppBuilder  IAppBuilder
		Rtr *mux.Router
}

func (handler *serverHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	handler.AppBuilder.InvokeMiddlewares(writer, request)

	handler.Rtr.ServeHTTP(writer, request)
}
