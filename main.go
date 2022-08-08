package main

import (
	"github.com/pquerna/ffjson/ffjson"
	"github.com/valyala/fasthttp"
)

type Ctx = fasthttp.RequestCtx

type Method string

type Routes map[string]finalRoute // string is the path of the route

type Header struct {
	Name  string
	Value string
}

type Middleware func(ctx *Ctx) bool

type Route struct {
	Path       string
	Method     Method
	Handle     func(ctx *fasthttp.RequestCtx)
	Children   []Route
	Error      func(ctx *fasthttp.RequestCtx)
	Middleware []Middleware
	Headers    []Header
}

type finalRoute struct {
	Method     Method
	Handle     func(ctx *fasthttp.RequestCtx)
	Middleware []Middleware
	Headers    []Header
	Error      func(ctx *fasthttp.RequestCtx)
}

type AppConfig struct {
	Routes     Routes
	Middleware []Middleware
	Headers    []Header
	Port       string
	IP         string
	Favicon    string
	NotFound   func(ctx *fasthttp.RequestCtx)
}

func JSON(data interface{}) []byte {

	json, err := ffjson.Marshal(data)

	if err != nil {
		panic(err)
	}

	return json

}

// func main() {

// 	App(
// 		AppConfig{
// 			Routes: CreateRouter([]Route{
// 				{
// 					Path:   "/",
// 					Method: "GET",
// 					Handle: func(ctx *Ctx) {
// 						ctx.WriteString("Hello World")
// 					},
// 				},
// 				{
// 					Path:   "/test/map",
// 					Method: "GET",
// 					Handle: func(ctx *Ctx) {
// 						ctx.Write(JSON(map[string]string{"test": "test"}))
// 					},
// 				},
// 				{
// 					Path:   "/static/user.json",
// 					Method: "GET",
// 					Handle: func(ctx *Ctx) {

// 						type user struct {
// 							Name string
// 							Age  int
// 						}

// 						ctx.Write(JSON(user{
// 							Name: "Piter",
// 							Age:  17,
// 						}))

// 					},
// 				},
// 				{
// 					Path:   "/static",
// 					Method: "GET",
// 					Handle: func(ctx *Ctx) {
// 						ctx.WriteString("Static Resources")
// 					},
// 					Middleware: []Middleware{
// 						func(ctx *Ctx) bool {
// 							ctx.Response.Header.Set("Content-Type", "text/html")
// 							return true
// 						},
// 					},
// 					Children: []Route{
// 						{
// 							Path:   "/index.html",
// 							Method: "GET",
// 							Handle: func(ctx *Ctx) {
// 								ctx.SendFile("index.html")
// 							},
// 							Middleware: []Middleware{
// 								func(ctx *Ctx) bool {
// 									ctx.Response.Header.Set("Static", "true")
// 									return true
// 								},
// 							},
// 						},
// 					},
// 				},
// 			}),
// 			Middleware: []Middleware{
// 				func(ctx *Ctx) bool { // Middleware passing the request to the next middleware or route returns true if the request is valid and false if it is not valid
// 					fmt.Println("Middleware")
// 					return true
// 				},
// 			},
// 			Headers: []Header{
// 				{
// 					Name:  "X-Powered-By",
// 					Value: "Alf",
// 				},
// 			},
// 			NotFound: func(ctx *Ctx) {
// 				ctx.WriteString("Route not found")
// 				ctx.SetStatusCode(fasthttp.StatusNotFound)
// 			},
// 		},
// 	)

// }

func (m Method) valid() bool {

	switch m {
	case "GET", "POST", "PUT", "DELETE":
		return true
	default:
		return false
	}

}

func createRoute(r finalRoute) finalRoute {

	if !r.Method.valid() {
		panic("Error Invalid method")
	}

	r.Error = func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("Route Error")
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
	}

	return r

}

func CreateRouter(r []Route) Routes {

	routes := make(Routes)

	for _, route := range r {

		if route.Path == "" {
			panic("Error Invalid path: " + route.Path)
		}

		if route.Path == "/" || (route.Children != nil && len(route.Children) > 0) { // if the route is a root route or has children
			for _, child := range route.Children {

				if child.Path == "" || child.Path == "/" {
					panic("Error Invalid path: " + route.Path + child.Path)
				}

				routes[route.Path+child.Path] = createRoute(finalRoute{
					Method:     child.Method,
					Handle:     child.Handle,
					Middleware: append(route.Middleware, child.Middleware...),
					Headers:    append(route.Headers, child.Headers...),
					Error:      child.Error,
				})
			}
		}

		routes[route.Path] = createRoute(finalRoute{
			Method:     route.Method,
			Handle:     route.Handle,
			Middleware: route.Middleware,
			Headers:    route.Headers,
			Error:      route.Error,
		})
	}

	return routes

}

func App(config AppConfig) {

	var r = config.Routes
	var h = config.Headers
	var m = config.Middleware

	var port string
	var ip string

	if config.Port == "" {
		port = "8080"
	} else {
		port = config.Port
	}

	if config.IP == "" {
		ip = "0.0.0.0"
	} else {
		ip = config.IP
	}

	println("Server running on " + ip + ":" + port)

	fasthttp.ListenAndServe(ip+":"+port, func(ctx *fasthttp.RequestCtx) {

		var method = string(ctx.Method())

		route, pathFound := r[string(ctx.Path())]

		if pathFound {

			for _, header := range h {
				ctx.Response.Header.Set(header.Name, header.Value)
			}

			if string(route.Method) == method {

				var next bool

				for _, middleware := range m {
					next = middleware(ctx) // if middleware returns false, it will stop the execution of the route
					if !next {
						break
					}
				}

				for _, middleware := range route.Middleware {

					next = middleware(ctx) // if middleware returns false, it will stop the execution of the route

					if !next {
						break
					}

				}

				if next {
					route.Handle(ctx)
				}

			} else {

				ctx.WriteString("Method not allowed " + method)

				route.Error(ctx)
			}

		} else {

			config.NotFound(ctx)

		}

	})

}
