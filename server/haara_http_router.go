package server

import (
	"fmt"
	"maps"
	"net/http"
)

type HaaraHttpMethod string

const (
	HttpMethodGET    HaaraHttpMethod = "GET"
	HttpMethodPOST   HaaraHttpMethod = "POST"
	HttpMethodPUT    HaaraHttpMethod = "PUT"
	HttpMethodPATCH  HaaraHttpMethod = "PATCH"
	HttpMethodDELETE HaaraHttpMethod = "DELETE"
)

type HaaraHttpRouter struct {
	handler *http.ServeMux
	routes  map[string]HaaraHttpHandlerFunc
}

func NewHaaraHttpRouter(handler *http.ServeMux) *HaaraHttpRouter {
	if handler == nil {
		handler = http.NewServeMux()
	}

	return &HaaraHttpRouter{
		handler: handler,
		routes:  make(map[string]HaaraHttpHandlerFunc),
	}
}

// method: "GET"|"POST"|any
// path: "/any/path"
func (hhr *HaaraHttpRouter) AddRoute(method HaaraHttpMethod, path string, handlerFunc HaaraHttpHandlerFunc) *HaaraHttpRouter {
	route := fmt.Sprintf("%s %s", method, path)
	hhr.routes[route] = handlerFunc
	hhr.handler.HandleFunc(route, handlerFunc)
	return hhr
}

func (hhr *HaaraHttpRouter) Get(path string, handlerFunc HaaraHttpHandlerFunc) *HaaraHttpRouter {
	return hhr.AddRoute(HttpMethodGET, path, handlerFunc)
}

func (hhr *HaaraHttpRouter) Post(path string, handlerFunc HaaraHttpHandlerFunc) *HaaraHttpRouter {
	return hhr.AddRoute(HttpMethodPOST, path, handlerFunc)
}

func (hhr *HaaraHttpRouter) Put(path string, handlerFunc HaaraHttpHandlerFunc) *HaaraHttpRouter {
	return hhr.AddRoute(HttpMethodPUT, path, handlerFunc)
}

func (hhr *HaaraHttpRouter) Patch(path string, handlerFunc HaaraHttpHandlerFunc) *HaaraHttpRouter {
	return hhr.AddRoute(HttpMethodPATCH, path, handlerFunc)
}

func (hhr *HaaraHttpRouter) Delete(path string, handlerFunc HaaraHttpHandlerFunc) *HaaraHttpRouter {
	return hhr.AddRoute(HttpMethodDELETE, path, handlerFunc)
}

func (hhr *HaaraHttpRouter) GetRoutes() (copyRoutes map[string]HaaraHttpHandlerFunc) {
	copyRoutes = make(map[string]HaaraHttpHandlerFunc)
	maps.Copy(copyRoutes, hhr.routes)
	return
}
