package alf

import (
	"github.com/pquerna/ffjson/ffjson"
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
	Path       string // path
	Method     Method // method (only one)
	Handle     func(ctx *fasthttp.RequestCtx) // func that handles the route
	Children   []Route // children of the route (if the route isnt the root route) [all childrens inherit parents Middlewares and Headers ]
	Error      func(ctx *fasthttp.RequestCtx) // func that handles errors of the route
	Middleware []Middleware // middlewares of the route
	Headers    []Header // headers
}

type finalRoute struct {
	Method     Method 
	Handle     func(ctx *fasthttp.RequestCtx)
	Middleware []Middleware
	Headers    []Header
	Error      func(ctx *fasthttp.RequestCtx)
}

type AppConfig struct {
	Routes     routes // routes of the app
	Middleware []Middleware // global middlewares
	Headers    []Header // global headers
	Port       string // port of the app
	Favicon    string // path to the favicon
	NotFound   func(ctx *fasthttp.RequestCtx) // func that handles NotFound requests
}

func JSON(data interface{}) []byte { // utility function to convert a struct or map to json([]byte)

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

func CreateRouter(r []Route) routes { // creates the routes of the app

	routes := make(routes)

	for _, route := range r {

		if route.Path == "" {
			panic("Error Invalid path: " + route.Path)
		}

		if route.Path != "/" || (route.Children != nil && len(route.Children) > 0) { // if the route is a root route or has children
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

func App(config AppConfig) error { // creates the app and starts it

	var r = config.Routes
	var h = config.Headers
	var m = config.Middleware

	var port string
	var ip string
	var favicon string

	if config.Port == "" {
		port = "8080"
	} else {
		port = config.Port
	}

	if config.Favicon == "" {
		favicon = "/favicon.ico"
	} else {
		favicon = config.Favicon
	}

	faviconHandler := fasthttp.FSHandler(favicon, 1)

	r["/favicon.ico"] = finalRoute{
		Method: "GET",
		Handle: faviconHandler,
	}

	println("Server running on " + ip + ":" + port)

	err := fasthttp.ListenAndServe(":"+port, func(ctx *fasthttp.RequestCtx) {

		

		var method = string(ctx.Method())

		route, pathFound := r[string(ctx.Path())]

		if pathFound {

			for _, header := range h {
				ctx.Response.Header.Set(header.Name, header.Value)
			}

			if string(route.Method) == method {

				var next bool = true

				for _, middleware := range m {
					next = middleware(ctx) // if middleware returns false, it will stop the execution of the route
					if !next {
						break
					}
				}

				if route.Middleware != nil {

					for _, middleware := range route.Middleware {

						next = middleware(ctx) // if middleware returns false, it will stop the execution of the route

						if !next {
							break
						}

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

	if err != nil {
		return err
	}

	return nil

}
