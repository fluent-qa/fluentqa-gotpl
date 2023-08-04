package main

import (
	"fmt"
	do2 "qfluent-go/backups/internal/do"
)

/**
 * Wheel
 */
type Wheel struct {
}

/**
 * Engine
 */
type Engine struct {
}

/**
 * Car
 */
type Car struct {
	Engine *Engine
	Wheels []*Wheel
}

func (c *Car) Start() {
	println("vroooom")
}

/**
 * Run example
 */
func main() {
	injector := do2.New()

	// provide wheels
	do2.ProvideNamedValue(injector, "wheel-1", &Wheel{})
	do2.ProvideNamedValue(injector, "wheel-2", &Wheel{})
	do2.ProvideNamedValue(injector, "wheel-3", &Wheel{})
	do2.ProvideNamedValue(injector, "wheel-4", &Wheel{})

	// provide car
	do2.Provide(injector, func(i *do2.Injector) (*Car, error) {
		car := Car{
			Engine: do2.MustInvoke[*Engine](i),
			Wheels: []*Wheel{
				do2.MustInvokeNamed[*Wheel](i, "wheel-1"),
				do2.MustInvokeNamed[*Wheel](i, "wheel-2"),
				do2.MustInvokeNamed[*Wheel](i, "wheel-3"),
				do2.MustInvokeNamed[*Wheel](i, "wheel-4"),
			},
		}

		return &car, nil
	})

	// provide engine
	do2.Provide(injector, func(i *do2.Injector) (*Engine, error) {
		return &Engine{}, nil
	})

	// start car
	car := do2.MustInvoke[*Car](injector)
	car.Start()
	fmt.Println(car.Wheels)
}
