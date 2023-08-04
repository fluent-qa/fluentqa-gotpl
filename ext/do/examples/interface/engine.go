package main

import (
	"qfluent-go/backups/internal/do"
)

type Engine interface{}

type engineImplem struct {
}

func NewEngine(i *do.Injector) (Engine, error) {
	return &engineImplem{}, nil
}
