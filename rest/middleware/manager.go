package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalMiddlewares []Middleware
}

func NewManager() *Manager {
	return &Manager{
		globalMiddlewares: make([]Middleware, 0),
	}
}

func (mngr *Manager) Use(mw ...Middleware) {
	mngr.globalMiddlewares = append(mngr.globalMiddlewares, mw...)
}
func (mngr *Manager) With(next http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {

		next = middleware(next)
	}

	return next

}

func (mngr *Manager) WrapMux(handler http.Handler) http.Handler {
	next := handler
	for _, globalMiddleware := range mngr.globalMiddlewares {
		next = globalMiddleware(next)

	}
	return next
}
