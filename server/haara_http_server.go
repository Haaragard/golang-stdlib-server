package server

import (
	"fmt"
	"net/http"
)

type HaaraHttpHandlerFunc func(http.ResponseWriter, *http.Request)

// host: "0.0.0.0"
// port: "8000"
// routes ["GET /health" => func(http.ResponseWriter, *http.Request)]
type HaaraHttpServer struct {
	host   string
	port   string
	srv    *http.Server
	mux    *http.ServeMux
	routes map[string]HaaraHttpHandlerFunc
}

func NewHaaraHttpServer() *HaaraHttpServer {
	haaraHttpServer := &HaaraHttpServer{
		host:   "0.0.0.0",
		port:   "8000",
		routes: make(map[string]HaaraHttpHandlerFunc),
	}

	haaraHttpServer.mux = http.NewServeMux()
	haaraHttpServer.srv = &http.Server{Handler: haaraHttpServer.mux}
	haaraHttpServer.updateSrvAddr()

	return haaraHttpServer
}

func (hhs *HaaraHttpServer) updateSrvAddr() {
	hhs.srv.Addr = fmt.Sprintf("%s:%s", hhs.host, hhs.port)
}

// host: "0.0.0.0"
func (hhs *HaaraHttpServer) SetHost(host string) *HaaraHttpServer {
	hhs.host = host
	hhs.updateSrvAddr()
	return hhs
}

// port: "8000"
func (hhs *HaaraHttpServer) SetPort(port string) *HaaraHttpServer {
	hhs.port = port
	hhs.updateSrvAddr()
	return hhs
}

// method: "GET"|"POST"|any
// path: "/any/path"
func (hhs *HaaraHttpServer) AddRoute(method, path string, handlerFunc HaaraHttpHandlerFunc) *HaaraHttpServer {
	route := fmt.Sprintf("%s %s", method, path)
	hhs.routes[route] = handlerFunc
	hhs.mux.HandleFunc(route, handlerFunc)
	return hhs
}

func (hhs *HaaraHttpServer) ServeAndListen() error {
	return hhs.srv.ListenAndServe()
}
