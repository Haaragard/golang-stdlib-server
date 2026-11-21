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
	server *http.Server
	router *HaaraHttpRouter
}

func NewHaaraHttpServer() *HaaraHttpServer {
	mux := http.NewServeMux()
	haaraHttpServer := &HaaraHttpServer{
		host:   "0.0.0.0",
		port:   "8000",
		server: &http.Server{Handler: mux},
		router: NewHaaraHttpRouter(mux),
	}

	haaraHttpServer.updateSrvAddr()

	return haaraHttpServer
}

func (hhs *HaaraHttpServer) updateSrvAddr() {
	hhs.server.Addr = fmt.Sprintf("%s:%s", hhs.host, hhs.port)
}

// host: "0.0.0.0"
func (hhs *HaaraHttpServer) SetHost(host string) *HaaraHttpServer {
	hhs.host = host
	hhs.updateSrvAddr()
	fmt.Printf("HaaraHttpServer host set to \"%s\"\n", host)
	return hhs
}

// port: "8000"
func (hhs *HaaraHttpServer) SetPort(port string) *HaaraHttpServer {
	hhs.port = port
	hhs.updateSrvAddr()
	fmt.Printf("HaaraHttpServer port set to \"%s\"\n", port)
	return hhs
}

func (hhs *HaaraHttpServer) GetRouter() *HaaraHttpRouter {
	return hhs.router
}

func (hhs *HaaraHttpServer) ServeAndListen() error {
	fmt.Printf("Server listening on %s:%s\n", hhs.host, hhs.port)
	return hhs.server.ListenAndServe()
}
