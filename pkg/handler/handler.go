package handler

import (
	"github.com/Alan15r/apiServiceCore/pkg/service"
	"github.com/fasthttp/router"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *router.Router {
	router := router.New()

	hellow := router.Group("/hellow")
	hellow.GET("/word/{user}", h.helloWord)

	return router
}
