package alf

import "github.com/valyala/fasthttp"

func handleMiddleware(ctx *fasthttp.RequestCtx, m []Middleware, next *bool) {

	for _, middleware := range m {
		*next = middleware(ctx) // if middleware returns false, it will stop the execution of the route
		if !*next {
			break
		}
	}

}
