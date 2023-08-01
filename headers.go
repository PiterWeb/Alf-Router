package alf

func handleHeaders(ctx *Ctx, h *[]Header) {

	for _, header := range *h {
		ctx.Response.Header.Set(header.Name, header.Value) // write global headers
	}

}
