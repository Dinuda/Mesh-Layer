package core

import (
	"net/http"
)

type Mesh struct {
	Function func(a MeshInput)
	Info     MeshInfo
}

type MeshInput struct {
	Re     *http.Request
	Strand *Strand
}

type MeshInfo struct {
	Name        string
	Description string
	Author      string
}

type MeshRegistry []Mesh

func (m *Mesh) Run(a MeshInput) {
	m.Function(a)
}
