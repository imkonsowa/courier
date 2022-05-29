package providers

import (
	"fmt"

	"courier/services/courier/app"
)

type Provider interface {
	Boot() error
}

func Ignite(a *app.App) {
	err := NewRouteProvider(a).Boot()
	if err != nil {
		panic(fmt.Sprintf("failed to ignite route provider; Err: %v", err))
	}

}
