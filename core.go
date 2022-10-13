package alf

import (
	"bytes"
	"github.com/pterm/pterm"
	"github.com/valyala/fasthttp"
)

func App(config AppConfig) error { // creates the app and starts it

	if config.Port == "" {
		config.Port = "8080"
	}

	if config.NotFound == nil {
		config.NotFound = func(ctx *fasthttp.RequestCtx) {
			ctx.WriteString("Path not found: ERROR 404")
		}
	}

	if config.NotAllowed == nil {
		config.NotAllowed = func(ctx *fasthttp.RequestCtx) {
			var method = string(ctx.Method())
			ctx.WriteString("Method (" + method + ") not allowed")
		}
	}

	pterm.Info.Println("Server running on  port :" + config.Port)

	err := fasthttp.ListenAndServe(":"+config.Port, func(ctx *fasthttp.RequestCtx) {

		handleRoute(ctx, config)

	})

	if err != nil {
		return err
	}

	return nil

}

func handleRoute(ctx *fasthttp.RequestCtx, config AppConfig) {

	routes, methodFound := config.Routes[string(ctx.Method())]

	if methodFound {

		route, pathFound := routes[string(ctx.Path())] // Intentar evitar la conversión de tipo del ctx.Path()

		if pathFound {

			handleHeaders(ctx, config.Headers)

			var next bool = true

			handleMiddleware(ctx, config.Middleware, &next) // handle global middleware

			if route.Middleware != nil && next {

				handleMiddleware(ctx, route.Middleware, &next) // handle specific middleware

			}

			if next {
				errRoute := route.Handle(ctx)

				if errRoute != nil {
					ctx.Response.Reset()
					route.Error(ctx, errRoute)

					showError("Route Error (" + errRoute.Error() + ") caused on route: " + string(ctx.Path()))

				}
			}

		} else if config.ServeStatic && (string(ctx.Method()) == "GET") {

			staticPrefix := []byte("/static/")
			staticHandler := fasthttp.FSHandler("/static", 1)

			if bytes.HasPrefix(ctx.Path(), staticPrefix) {

				staticHandler(ctx)
			} else {

				config.NotFound(ctx)
				showWarning("Route not Found: " + string(ctx.Path()))

			}

		} else {

			config.NotFound(ctx)
			showWarning("Route not Found: " + string(ctx.Path()))

		}

	} else {

		config.NotAllowed(ctx)

	}

}
