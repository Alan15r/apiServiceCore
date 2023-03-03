package handler

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

func (h *Handler) helloWord(ctx *fasthttp.RequestCtx) {
	user := ctx.UserValue("user").(string)
	if err := h.services.ServiceHellow(user); err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	ctx.SetBodyString(fmt.Sprintf("Hellow, %s", user))
	ctx.SetStatusCode(fasthttp.StatusOK)
}
