package main

import (
	"fmt"

	"github.com/awoodbeck/plugins/server"
)

type greeter struct{}

func (g greeter) Name() string {
	return "salut"
}

func (g greeter) Greet() {
	fmt.Println("Salut!")
}

func Salut() server.Greeter {
	return new(greeter)
}
