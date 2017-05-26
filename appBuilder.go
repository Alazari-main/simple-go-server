package server 

import (
	"net/http"
	"server/middleware"
)

var (
	builder *appBuilder
)

// IAppBuilder provides methods for registering middlewares into pipeline
type IAppBuilder interface {
	UseMiddleware(middleware middleware.IMiddleware)

	UseMiddlewareFunc(middleware.MiddlewareFunc)

	InvokeMiddlewares(http.ResponseWriter, *http.Request)
}

type appBuilder struct {
	middlewares []middleware.IMiddleware
}

// UseMiddleware appends specified middleware into execution pipeline.
func (app *appBuilder) UseMiddleware(middleware middleware.IMiddleware) {
	app.middlewares = append(app.middlewares, middleware)
}

// UseFunc appends middleware function into execution pipeline.
func (app *appBuilder) UseMiddlewareFunc(middlewareFunc middleware.MiddlewareFunc) {
	mdw := middleware.NewAnonymousMiddleware(middlewareFunc)

	app.middlewares = append(app.middlewares, mdw)
}

func (app *appBuilder) InvokeMiddlewares(writer http.ResponseWriter, request *http.Request) {
	if len(app.middlewares) == 0 {
		return
	}

	context := middleware.NewMiddlewareContext(writer, request)

	for _, middleware := range app.middlewares {
		middleware.Apply(context)
	}
}

// NewAppBuilder returns new IAppBuilder
func NewAppBuilder() IAppBuilder {
	if builder == nil {
		builder = new(appBuilder)
	}

	builder.middlewares = make([]middleware.IMiddleware, 0)

	return builder
}
