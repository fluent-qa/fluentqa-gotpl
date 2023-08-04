package main

import (
	do2 "qfluent-go/backups/internal/do"
)

type Car interface {
	Start()
}

type carImplem struct {
	Engine Engine
	Wheels []*Wheel
}

func (c *carImplem) Start() {
	println("vroooom")
}

func NewCar(i *do2.Injector) (Car, error) {
	wheels := []*Wheel{
		do2.MustInvokeNamed[*Wheel](i, "wheel-1"),
		do2.MustInvokeNamed[*Wheel](i, "wheel-2"),
		do2.MustInvokeNamed[*Wheel](i, "wheel-3"),
		do2.MustInvokeNamed[*Wheel](i, "wheel-4"),
	}

	engine := do2.MustInvoke[Engine](i)

	car := carImplem{
		Engine: engine,
		Wheels: wheels,
	}

	return &car, nil
}
