package core

import (
	"fmt"
	"html"
	"net/http"
)

type Cloth []Fabric

// returns a http.HandlerFunc
func (c *Cloth) Handle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handled := false
		killed := false
		for i := 0; i < len(*c); i++ {
			uC := *c
			if html.EscapeString(r.URL.Path) == uC[i].Path {
				s := Strand{Parent: &uC[i], Rw: &w, R: r, Killed: &killed}
				for i := uint(0); i < s.GetMeshCount(); i++ {
					s.NextMesh()
				}
				if !*s.Killed {
					s.E()
				}
				handled = true
				break
			}
		}
		if !handled {
			_, _ = fmt.Fprintln(w, []byte("Could not find path"))
		}
	}
}
