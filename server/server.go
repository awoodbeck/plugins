package server

import "sync"

var (
	mu       sync.Mutex
	greeters = make(map[string]Greeter)
)

func RegisterGreeter(g Greeter) {
	if g != nil {
		mu.Lock()
		greeters[g.Name()] = g
		mu.Unlock()
	}
}

func Greet() {
	mu.Lock()
	defer mu.Unlock()

	for _, g := range greeters {
		g.Greet()
	}
}

type Greeter interface {
	Greet()
	Name() string
}
