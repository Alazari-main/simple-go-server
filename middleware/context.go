package middleware

import (
	"net/http"
)

// MiddlewareContext contains useful data srtuctures for middlewares.
type MiddlewareContext struct {
	Request        *http.Request
	ResponceWriter http.ResponseWriter
}

// NewMiddlewareContext creates new middlewareContext
func NewMiddlewareContext(writer http.ResponseWriter, request *http.Request) *MiddlewareContext {
	context := &MiddlewareContext{Request: request, ResponceWriter: writer}

	return context
}
