package main

import (
	do2 "qfluent-go/backups/internal/do"
)

func main() {
	injector := do2.New()

	// provide wheels
	do2.ProvideNamedValue(injector, "wheel-1", NewWheel())
	do2.ProvideNamedValue(injector, "wheel-2", NewWheel())
	do2.ProvideNamedValue(injector, "wheel-3", NewWheel())
	do2.ProvideNamedValue(injector, "wheel-4", NewWheel())

	// provide car
	do2.Provide(injector, NewCar)

	// provide engine
	do2.Provide(injector, NewEngine)

	// start car
	car := do2.MustInvoke[Car](injector)
	car.Start()
}
