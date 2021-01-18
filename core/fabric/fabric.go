package fabric

import (
	"fmt"
	"io"
	"log"
	"mesh/core/mesh"
	"mesh/core/server"
	"net/http"
)

type Fabric struct {
	Meshes    mesh.Registry
	Path      string
	RouteDest server.API
	uid       string
	meshCount uint
}

type Strand struct {
	currentMesh uint
	Parent      *Fabric
	Rw          *http.ResponseWriter
	R           *http.Request
}

func (s *Strand) NextMesh() {
	if s.currentMesh != s.Parent.meshCount-1 {
		rMesh := s.Parent.Meshes[s.currentMesh]
		rMesh.Run(mesh.Input{Re: s.R})
		s.currentMesh++
	}
}

func (s *Strand) E() {
	e, r := s.Parent.RouteDest.Send(mesh.Info{})
	if e != nil {
		log.Fatalln(e)
	}
	_, err := io.Copy(*s.Rw, r.Body)
	if err != nil {
		_, _ = fmt.Fprintln(*s.Rw, []byte("An error occurred"))
	}
}

func (s Strand) GetMeshCount() uint {
	return uint(len(s.Parent.Meshes))
}
