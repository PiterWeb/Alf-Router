package alf

import (
	// "fmt"
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
	"github.com/valyala/fasthttp"
	"sync"
	// "time"
)

var wg sync.WaitGroup

func createRoute(r finalRoute) finalRoute { // create the route with the given parameters

	if r.Error == nil {

		r.Error = func(ctx *fasthttp.RequestCtx, err error) {
			ctx.WriteString("Route Error: " + err.Error())
			ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		}

	}

	return r

}

func createChildrenRoutes(routes methodRoutes, r Route, initialPath string) {

	if r.Path != "/" && (r.Children != nil && len(r.Children) > 0) { // if the route is not the root route and has children

		initialPath = initialPath + r.Path

		for _, child := range r.Children {

			// start := time.Now()

			if !child.Method.valid() {
				showError("Invalid method ( " + child.Method.string() + " ) on route " + initialPath + child.Path)
			}

			if child.Path == "" || child.Path == "/" {
				showInternalError("Invalid path set on route: ( " + initialPath + child.Path + " )")
			}

			wg.Add(1)

			go func() {

				routes[r.Method.string()][initialPath+child.Path] = createRoute(finalRoute{ // generate new subroute
					Method:     child.Method,
					Handle:     child.Handle,
					Middleware: append(r.Middleware, child.Middleware...), // apply middlewares of the parent route
					Headers:    append(r.Headers, child.Headers...),       // apply headers of the parent route
					Error:      child.Error,
				})

				wg.Done()

			}()

			if child.Children != nil && len(child.Children) > 0 {
				createChildrenRoutes(routes, child, initialPath)
			}

			// fmt.Printf("Child Route: (%s) => %.10d\n", initialPath+child.Path, time.Since(start).Nanoseconds())

		}

	}

}

func CreateRouter(r []Route) methodRoutes { // creates the routes of the app

	pterm.DefaultBigText.WithLetters(putils.LettersFromStringWithStyle("ALF", pterm.NewStyle(pterm.FgBlue))).Render()

	routes := methodRoutes{
		"GET":    make(map[string]finalRoute),
		"POST":   make(map[string]finalRoute),
		"DELETE": make(map[string]finalRoute),
		"PUT":    make(map[string]finalRoute),
	}

	for _, route := range r {

		if !route.Method.valid() {
			showError("Invalid method ( " + route.Method.string() + " ) on route " + route.Path)
		}

		if route.Path == "" {
			showInternalError("Invalid path set on route: (" + route.Path + " )")
		}

		createChildrenRoutes(routes, route, "")

		// start := time.Now()

		go func(route Route) {

			routes[route.Method.string()][route.Path] = createRoute(finalRoute{
				Method:     route.Method,
				Handle:     route.Handle,
				Middleware: route.Middleware,
				Headers:    route.Headers,
				Error:      route.Error,
			})

			// fmt.Printf("Normal Route: (%s) => %.10d\n", route.Path, time.Since(start).Nanoseconds())

		}(route)

		wg.Wait()
	}

	return routes

}
