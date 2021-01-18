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
	Killed      *bool
}

func (s *Strand) NextMesh() {
	if s.currentMesh != s.Parent.meshCount-1 && !*s.Killed {
		rMesh := s.Parent.Meshes[s.currentMesh]
		rMesh.Run(mesh.Input{Re: s.R, Kill: s.Killed})
		s.currentMesh++
	}
}

func (s *Strand) E() {
	e, r := s.Parent.RouteDest.Send()
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
