package alf

import (
	"sync"

	"github.com/valyala/fasthttp"
)

type Ctx = fasthttp.RequestCtx // alias for Ctx

type Method string // GET, POST, PUT, DELETE

type Header struct {
	Name  string // name of the header
	Value string // value of the header
}

type Middleware func(ctx *Ctx) bool // func that returns true if passed and false if an error ocurred

type Route struct {
	Path       string                    // path
	Method     Method                    // method (only one)
	Handle     func(ctx *Ctx) error      // func that handles the route
	Children   []Route                   // children of the route (if the route isnt the root route) [all childrens inherit parents Middlewares and Headers ]
	Error      func(ctx *Ctx, err error) // func that handles errors of the route (if err returned on handle method will be invoked with the error in the parameters)
	Middleware []Middleware              // middlewares of the route
	Headers    []Header                  // headers
}

type finalRoute struct {
	Method     Method
	Handle     func(ctx *Ctx) error
	Middleware []Middleware
	Headers    []Header
	Error      func(ctx *Ctx, err error)
}

type AppConfig struct {
	Routes      methodRoutes   // routes of the app
	Middleware  []Middleware   // global middlewares
	Headers     []Header       // global headers
	Port        string         // port of the app | default value '8080'
	NotFound    func(ctx *Ctx) // func that handles NotFound requests
	ServeStatic bool           // if true, the app will serve static files on "/static"
	BeforeInit  func(*AppConfig)
}

type routes struct {
	mu sync.Mutex
	methodRoutes
}

type methodRoutes map[string]map[string]finalRoute // 1º string = method ("GET", ...) ; 2º string = path ("/api", ...) ;
