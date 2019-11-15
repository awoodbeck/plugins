package main

import (
	"fmt"

	"github.com/awoodbeck/plugins/server"
)

type greeter struct{}

func (g greeter) Name() string {
	return "hello"
}

func (g greeter) Greet() {
	fmt.Println("Hello!")
}

func Hello() server.Greeter {
	return new(greeter)
}
