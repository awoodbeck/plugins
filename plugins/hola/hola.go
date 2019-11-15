package main

import (
	"fmt"

	"github.com/awoodbeck/plugins/server"
)

type greeter struct{}

func (g greeter) Name() string {
	return "hola"
}

func (g greeter) Greet() {
	fmt.Println("Hola!")
}

func Hola() server.Greeter {
	return new(greeter)
}
