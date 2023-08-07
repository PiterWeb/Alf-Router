package alf

import (
	"bytes"
	"fmt"

	misc "github.com/PiterWeb/Alf-Router/errors"
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
	"github.com/valyala/fasthttp"
)

func App(config *AppConfig) error { // creates the app and starts it

	if config.Port == "" {
		config.Port = "8080"
	}

	if config.NotFound == nil {
		config.NotFound = func(ctx *Ctx) {
			ctx.Response.SetStatusCode(404)
			ctx.WriteString("Path not found: ERROR 404")
		}
	}

	for _, pl := range config.Plugins {

		err := pl.Init_plugin()

		if err != nil {
			misc.ShowError(err.Error())
		}

	}

	pterm.DefaultBigText.WithLetters(putils.LettersFromStringWithStyle("ALF", pterm.NewStyle(pterm.FgBlue))).Render()

	pterm.Info.Println("Server running on  port :" + config.Port)

	errChan := make(chan error)

	go func() {
		errChan <- fasthttp.ListenAndServe(":"+config.Port, func(ctx *Ctx) {

			handleRoute(ctx, config)

		})
	}()

	go func() {
		pterm.Info.Println("Press enter to stop the server: ")

		fmt.Scanf("\n%c")

		errChan <- nil
	}()

	return <-errChan

}

func handleRoute(ctx *Ctx, config *AppConfig) {

	routes, methodFound := config.Routes[string(ctx.Method())]

	if !methodFound {
		return
	}

	var path string = string(ctx.Path())

	if len(path) > 1 && path[len(path)-1] == '/' { // make /api equal to /api/

		path = path[:len(path)-1]

	}

	route, pathFound := routes[path] // Intentar evitar la conversión de tipo del ctx.Path()

	if !pathFound {

		config.NotFound(ctx)
		go misc.ShowWarning("Route not Found: " + string(ctx.Path()) + " [" + string(ctx.Method()) + "]")
		return

	}

	if config.ServeStatic && (string(ctx.Method()) == "GET") {

		staticPrefix := []byte("/static/")
		staticHandler := fasthttp.FSHandler("/static", 1)

		if bytes.HasPrefix(ctx.Path(), staticPrefix) {
			staticHandler(ctx)
			return
		}

	}

	handleHeaders(ctx, &config.Headers)

	next := true

	handleMiddleware(ctx, &config.Middleware, &next) // handle global middleware

	if route.Middleware != nil && next {

		handleMiddleware(ctx, &route.Middleware, &next) // handle specific middleware

	}

	var errRoute error

	if next {
		errRoute = route.Handle(ctx)
	}

	if errRoute != nil {
		ctx.Response.Reset()
		route.Error(ctx, errRoute)
		go misc.ShowError("Route Error (" + errRoute.Error() + ") caused on route: " + string(ctx.Path()))
		return
	}

}
