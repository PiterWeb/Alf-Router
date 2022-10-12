package alf

import (
	"bytes"
	"github.com/pterm/pterm"
	"github.com/valyala/fasthttp"
)

func App(config AppConfig) error { // creates the app and starts it

	var r = config.Routes
	var h = config.Headers
	var m = config.Middleware

	var port string
	var favicon string

	if config.Port == "" {
		port = "8080"
	} else {
		port = config.Port
	}

	staticPrefix := []byte("/static/")
	staticHandler := fasthttp.FSHandler("/static", 1)

	faviconHandler := fasthttp.FSHandler(favicon, 1)

	r["/favicon.ico"] = finalRoute{
		Method: "GET",
		Handle: faviconHandler,
	}

	pterm.Info.Println("Server running on  port :" + port)

	err := fasthttp.ListenAndServe(":"+port, func(ctx *fasthttp.RequestCtx) {

		var method = string(ctx.Method())

		route, pathFound := r[string(ctx.Path())]

		if pathFound {

			for _, header := range h {
				ctx.Response.Header.Set(header.Name, header.Value)
			}

			if string(route.Method.string()) == method {

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

				if route.Error != nil {
					route.Error(ctx)
				}

			}

		} else if config.ServeStatic {

			if bytes.HasPrefix(ctx.Path(), staticPrefix) {
				staticHandler(ctx)
			} else if config.NotFound != nil {
				config.NotFound(ctx)
				showWarning("Route not Found: " + string(ctx.Path()))

			} else {
				ctx.WriteString("Route not found: ERROR 404")
				showWarning("Route not Found: " + string(ctx.Path()))

			}

		} else {

			if config.NotFound != nil {
				config.NotFound(ctx)

			} else {
				ctx.WriteString("Route not found: ERROR 404")
			}

			showWarning("Route not Found: " + string(ctx.Path()))

		}

	})

	if err != nil {
		return err
	}

	return nil

}
