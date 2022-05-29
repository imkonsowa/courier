package main

import (
	"courier/services/courier/app"
	"courier/services/courier/providers"
)

func main() {
	a := app.NewApp()

	providers.Ignite(a)

	a.Run()
}
