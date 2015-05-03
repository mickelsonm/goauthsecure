package middleware

import (
	"net/http"
)

var middleman *middlewareHandler

type (
	middlewareHandler struct{}
	Config            struct {
		RequireAuth bool
	}
)

func init() {
	middleman = new(middlewareHandler)
}

//this acts like a catch all for all routes
func Middleware(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	//TODO: could put some really neat stuff in here
	next(w, req)
}

//targeting a specific route
func Route(handler http.HandlerFunc, config *Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		middleman.ServeHTTP(w, r, handler, config)
	}
}

//statisfies the http.Handler requirement
func (mh *middlewareHandler) ServeHTTP(w http.ResponseWriter, req *http.Request, next http.Handler, config *Config) {
	if config != nil {
		//TODO: we could implement more in here if we wanted to
		if config.RequireAuth {
			//TODO: implement all the authorization checking/etc
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
	}
	//proceeds to the next/destination handler
	next.ServeHTTP(w, req)
}
