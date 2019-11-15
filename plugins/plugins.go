//go:generate go build -buildmode=plugin -o hello/hello.so hello/hello.go
//go:generate go build -buildmode=plugin -o hola/hola.so hola/hola.go
//go:generate go build -buildmode=plugin -o salut/salut.so salut/salut.go

package plugins

var plugins = []struct {
	name   string
	symbol string
}{
	{name: "hello", symbol: "Hello"},
	{name: "hola", symbol: "Hola"},
	{name: "salut", symbol: "Salut"},
}
