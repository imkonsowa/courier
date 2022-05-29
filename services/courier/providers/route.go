package providers

import (
	"courier/services/courier/app"
	"courier/services/courier/data/adapters"
	"courier/services/courier/http/handlers"
	"courier/services/courier/http/router"
)

type RouteProvider struct {
	App *app.App
}

func NewRouteProvider(a *app.App) *RouteProvider {
	return &RouteProvider{
		App: a,
	}
}

func (r *RouteProvider) Boot() error {
	parcelHandler := handlers.NewParcelHandler(
		adapters.NewParcelAdapter(r.App.DB),
	)

	rtr := &router.Router{
		App: r.App,

		ParcelHandler: parcelHandler,
	}

	rtr.RegisterAPIRoutes()

	return nil
}
