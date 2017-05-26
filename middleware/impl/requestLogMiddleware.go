package impl

import (
	"server/middleware"
)

// RequestLogMiddleware provides request logging mechanism to simpleserver.
type RequestLogMiddleware struct {}

// Apply function is being invoked on each request and logs request information into specific file
func (middleware *RequestLogMiddleware) Apply(context *middleware.MiddlewareContext) {
	//TODO: Implement logging into file here
}