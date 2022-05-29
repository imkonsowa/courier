package router

import (
	"courier/services/courier/app"
	"courier/services/courier/http/handlers"
)

type Router struct {
	App *app.App

	ParcelHandler *handlers.ParcelHandler
}

func (r *Router) RegisterAPIRoutes() *Router {
	engine := r.App.Engine

	engine.GET("/cargo-file", r.ParcelHandler.GenerateCargoReport)

	return r
}
