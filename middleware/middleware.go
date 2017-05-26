package middleware

// IMiddleware is an interface for middlewares which will be invoked in chain.
type IMiddleware interface {
	Apply(context *MiddlewareContext)
}
