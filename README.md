# stdlib-rest

This project serves the purpose of learning golang stdlib while building a simple wrapper for stlib http server.

## Package structure

### container

A service register container for registering services for reusability.

### server

Has all the http wrapper essentials as Server, Router, own handler function interface.

#### HaaraHttpServer

Contains essential data for server definition and handling like `host: "0.0.0.0"`, `port: "8000"`, `http.Server` instance and more.

#### HaaraHttpHandlerFunc

Serves as an interface for Handler Function definition

#### Haara Http Router

Serves as an wrapper for route definition.

### controller

Has the controllers(HandlerFuncs) to deal with http requests

## main

Just test and make it works for simple uses.
