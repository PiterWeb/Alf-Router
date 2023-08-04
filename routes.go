package alf

import (
	"strings"
	"sync"

	misc "github.com/PiterWeb/Alf-Router/errors"

	"github.com/valyala/fasthttp"
)

var wg sync.WaitGroup

func (r *routes) addRoute(method string, path string, route *finalRoute) {

	r.mu.Lock()

	r.methodRoutes[method][path] = createRoute(&finalRoute{
		Method:     route.Method,
		Handle:     route.Handle,
		Middleware: route.Middleware,
		Headers:    route.Headers,
		Error:      route.Error,
	})
	r.mu.Unlock()

}

func createRoute(r *finalRoute) finalRoute { // create the route with the given parameters

	if r.Error == nil {

		r.Error = func(ctx *Ctx, err error) {
			ctx.WriteString("Route Error: " + err.Error())
			ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		}

	}

	return *r

}

func routeIsNotRootAndHaveChildren(r Route) bool {

	return r.Path != "/" && (r.Children != nil && len(r.Children) > 0)
}

func createChildrenRoutes(routes *routes, r Route, initialPath string) {

	if !routeIsNotRootAndHaveChildren(r) {
		return
	}

	newPath := initialPath + r.Path

	for _, child := range r.Children {

		child.Path = strings.Trim(child.Path, " ")

		newChildPath := newPath + child.Path
		// println("Full path: " + newChildPath)


		if !child.Method.valid() {
			misc.ShowError("Invalid method ( " + child.Method.string() + " ) on route " + newChildPath)
			continue
		}

		if child.Path == "" || child.Path == "/" {
			misc.ShowInternalError("Invalid path set on route: ( " + newChildPath + " )")
			continue
		}

		wg.Add(1)

		go func() {

			routes.addRoute(child.Method.string(), newChildPath, &finalRoute{ // generate new subroute
				Method:     child.Method,
				Handle:     child.Handle,
				Middleware: append(r.Middleware, child.Middleware...), // apply middlewares of the parent route
				Headers:    append(r.Headers, child.Headers...),       // apply headers of the parent route
				Error:      child.Error,
			})

			wg.Done()

		}()

		if child.Children != nil && len(child.Children) > 0 {
			createChildrenRoutes(routes, child, newPath)
		}

	}

}

func CreateRouter(r []Route) methodRoutes { // creates the routes of the app

	const initialPath string = ""

	routes := routes{
		methodRoutes: methodRoutes{
			"GET":     make(map[string]finalRoute),
			"POST":    make(map[string]finalRoute),
			"DELETE":  make(map[string]finalRoute),
			"PUT":     make(map[string]finalRoute),
			"PATCH":   make(map[string]finalRoute),
			"HEAD":    make(map[string]finalRoute),
			"OPTIONS": make(map[string]finalRoute),
		},
	}

	for _, route := range r {

		if !route.Method.valid() {
			misc.ShowError("Invalid method ( " + route.Method.string() + " ) on route " + route.Path)
			continue
		}

		if route.Path == "" {
			misc.ShowInternalError("Invalid path set on route: (" + route.Path + " )")
			continue
		}

		route.Path = strings.Trim(route.Path, " ")

		createChildrenRoutes(&routes, route, initialPath)

		wg.Add(1)

		go func(route Route) {

			routes.addRoute(route.Method.string(), route.Path, &finalRoute{
				Method:     route.Method,
				Handle:     route.Handle,
				Middleware: route.Middleware,
				Headers:    route.Headers,
				Error:      route.Error,
			})

			wg.Done()

		}(route)

		wg.Wait()
	}

	return routes.methodRoutes

}
