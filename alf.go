package alf

// Tareas:
// Intentar crear las rutas con concurrency para mejorar rendimiento inicial
// ORM =>

import (
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
	"github.com/valyala/fasthttp"
)

// Example APP:

// func main() {

// 	alf.App(
// 		alf.AppConfig{
// 			Routes: CreateRouter([]alf.Route{
// 				{
// 					Path:   "/",
// 					Method: "GET",
// 					Handle: func(ctx *Ctx) {
// 						ctx.WriteString("Hello World")
// 					},
// 				},
// 			}),
// 			Headers: []alf.Header{
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

func createRoute(r finalRoute) finalRoute { // create the route with the given parameters

	r.Error = func(ctx *fasthttp.RequestCtx, err error) {
		ctx.WriteString("Route Error: " + err.Error())
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
	}

	return r

}

func CreateRouter(r []Route) methodRoutes { // creates the routes of the app

	pterm.DefaultBigText.WithLetters(putils.LettersFromStringWithStyle("ALF", pterm.NewStyle(pterm.FgBlue))).Render()

	routes := map[string]map[string]finalRoute{
		"GET":    {},
		"POST":   {},
		"DELETE": {},
		"PUT":    {},
	}

	for _, route := range r {

		if !route.Method.valid() {
			showError("Invalid method ( " + route.Method.string() + " ) on route")
		}

		if route.Path == "" {
			showInternalError("Invalid path set on route: (" + route.Path + " )")
		}

		if route.Path != "/" && (route.Children != nil && len(route.Children) > 0) { // if the route is not the root route and has children
			for _, child := range route.Children {

				if child.Path == "" || child.Path == "/" {
					showInternalError("Invalid path set on route: ( " + route.Path + child.Path + " )")
				}

				routes[route.Method.string()][route.Path+child.Path] = createRoute(finalRoute{ // generate new subroute
					Method:     child.Method,
					Handle:     child.Handle,
					Middleware: append(route.Middleware, child.Middleware...), // apply middlewares of the parent route
					Headers:    append(route.Headers, child.Headers...),       // apply headers of the parent route
					Error:      child.Error,
				})
			}

		}

		routes[route.Method.string()][route.Path] = createRoute(finalRoute{
			Method:     route.Method,
			Handle:     route.Handle,
			Middleware: route.Middleware,
			Headers:    route.Headers,
			Error:      route.Error,
		})
	}

	return routes

}
