package mesh

import "net/http"

type Mesh struct {
	Function func(a Input)
	Info     Info
}

type Input struct {
	Re   *http.Request
	Kill *bool
}

type Info struct {
	Name        string
	Description string
	Author      string
}

type Registry []Mesh

func (m *Mesh) Run(a Input) {
	m.Function(a)
}
