package routes

import (
	mux "github.com/julienschmidt/httprouter"
)

type Route struct {
	Name   string
	Method string
	Path   string
	Handle mux.Handle
}

type Routes []Route

var routes = Routes{
	Route{
		"Show",
		"GET",
		"/store/{key}",
		Get,
	},
	Route{
		"Insert or Update",
		"PUT",
		"/store/{key}/{value}",
		Put,
	},
}
