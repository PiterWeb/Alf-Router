package alf

import (
	"strings"
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

func (m Method) valid() bool { // return if the method is valid

	switch m.string() {
	case "GET", "POST", "PUT", "DELETE":
		return true
	default:
		return false
	}

}

func (m Method) string() string {

	return strings.ToUpper(string(m))

}

func createRoute(r finalRoute) finalRoute { // create the route with the given parameters

	if !r.Method.valid() {
		showWarning("Invalid method ( " + string(r.Method) + " ) on route")
	}

	r.Error = func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("Route Error")
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		showError("Route Error caused on route: " + string(ctx.Path()))
	}

	return r

}

func CreateRouter(r []Route) routes { // creates the routes of the app

	pterm.DefaultBigText.WithLetters(putils.LettersFromStringWithStyle("ALF", pterm.NewStyle(pterm.FgBlue))).Render()

	routes := make(routes)

	for _, route := range r {

		if route.Path == "" {
			showInternalError("Invalid path set on route: ( " + route.Path + " )")
		}

		if route.Path != "/" && (route.Children != nil && len(route.Children) > 0) { // if the route is not the root route and has children
			for _, child := range route.Children {

				if child.Path == "" || child.Path == "/" {
					panic("Error Invalid path: " + route.Path + child.Path)
				}

				routes[route.Path+child.Path] = createRoute(finalRoute{ // generate new subroute
					Method:     child.Method,
					Handle:     child.Handle,
					Middleware: append(route.Middleware, child.Middleware...), // apply middlewares of the parent route
					Headers:    append(route.Headers, child.Headers...),       // apply headers of the parent route
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
