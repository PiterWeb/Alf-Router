package alf

import (
	"github.com/valyala/fasthttp"
)

type Ctx = fasthttp.RequestCtx // alias for fasthttp.RequestCtx

type Method string // GET, POST, PUT, DELETE

type routes map[string]finalRoute // string is the path of the route

type Header struct {
	Name  string // name of the header
	Value string // value of the header
}

type Middleware func(ctx *Ctx) bool // func that returns true if passed and false if an error ocurred

type Route struct {
	Path       string                         // path
	Method     Method                         // method (only one)
	Handle     func(ctx *fasthttp.RequestCtx) // func that handles the route
	Children   []Route                        // children of the route (if the route isnt the root route) [all childrens inherit parents Middlewares and Headers ]
	Error      func(ctx *fasthttp.RequestCtx) // func that handles errors of the route
	Middleware []Middleware                   // middlewares of the route
	Headers    []Header                       // headers
}

type finalRoute struct {
	Method     Method
	Handle     func(ctx *fasthttp.RequestCtx)
	Middleware []Middleware
	Headers    []Header
	Error      func(ctx *fasthttp.RequestCtx)
}

type AppConfig struct {
	Routes      routes                         // routes of the app
	Middleware  []Middleware                   // global middlewares
	Headers     []Header                       // global headers
	Port        string                         // port of the app
	NotFound    func(ctx *fasthttp.RequestCtx) // func that handles NotFound requests
	ServeStatic bool                           // if true, the app will serve static files on "/static"
}