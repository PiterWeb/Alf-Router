package alf

import "github.com/valyala/fasthttp"

func handleHeaders(ctx *fasthttp.RequestCtx, h []Header) {

	for _, header := range h {
		ctx.Response.Header.Set(header.Name, header.Value) // write global headers
	}

}
