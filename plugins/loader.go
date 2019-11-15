package plugins

import (
	"errors"
	"fmt"
	"path/filepath"
	"plugin"
	"runtime"

	"github.com/awoodbeck/plugins/server"
)

var (
	ErrInvalidGreeter = errors.New("invalid Greeter")

	self string
)

func init() {
	_, self, _, _ = runtime.Caller(0)

	registerPlugins()
}

func registerPlugins() {
	for _, p := range plugins {
		g, err := loadPlugin(filepath.Join(filepath.Dir(self), p.name, fmt.Sprintf("%s.so", p.name)), p.symbol)
		if err != nil {
			fmt.Println(err)
		}

		server.RegisterGreeter(g)
	}
}

func loadPlugin(path, symbol string) (server.Greeter, error) {
	p, err := plugin.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := p.Lookup(symbol)
	if err != nil {
		return nil, err
	}

	g, ok := s.(func() server.Greeter)
	if !ok {
		return nil, ErrInvalidGreeter
	}

	return g(), nil
}
