package middleware

// AnonymousMiddleware represents a function which can be used as a middleware.
type anonymousMiddleware struct {
	apply MiddlewareFunc
}

// Apply applies AnonymousMiddleware on specified IMiddlewareContext.
func (middleware *anonymousMiddleware) Apply(context *MiddlewareContext) {
	middleware.apply(context)
}

// NewAnonymousMiddleware creates new AnonymousMiddleware.
func NewAnonymousMiddleware(middlewareFunc MiddlewareFunc) IMiddleware {
	middleware := &anonymousMiddleware{apply: middlewareFunc}

	return middleware
}
