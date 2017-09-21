package server

import (
	"fmt"
	"net/http"
)

type RequestHandler func(http.ResponseWriter, *http.Request) bool

type RequestBroker struct {
	Routers                map[string][]RequestHandler
	HttpMethodNotSupported RequestHandler
	NoRouterServedRequest  RequestHandler
}

func (r *RequestBroker) serveOrRejectWithRouters(
	routers []RequestHandler, w http.ResponseWriter, req *http.Request) bool {
	for _, router := range routers {
		requestServed := router(w, req)
		if requestServed {
			return true
		}
	}

	return false
}

func (r *RequestBroker) Resolve(w http.ResponseWriter, req *http.Request) {
	routers, ok := r.Routers[req.Method]

	if ok {
		requestServed := r.serveOrRejectWithRouters(routers, w, req)
		if !requestServed {
			r.NoRouterServedRequest(w, req)
		}
	} else {
		r.HttpMethodNotSupported(w, req)
	}
}

func DefaultHandleUnsupportedHttpMethod(
	w http.ResponseWriter, req *http.Request) bool {
	fmt.Fprintf(w, "nope\nThis http method is unsupported: %s", req.Method)
	return true
}

func DefaultHandleNoRouterServedRequest(
	w http.ResponseWriter, req *http.Request) bool {
	fmt.Fprintf(w, "nope\nNone of handlers served request!")
	return true
}

func NewRequestBroker(routers map[string][]RequestHandler) *RequestBroker {
	return &RequestBroker{
		routers,
		DefaultHandleUnsupportedHttpMethod,
		DefaultHandleNoRouterServedRequest}
}
