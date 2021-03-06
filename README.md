# MESH LAYER

## What is Mesh Layer
mesh layer is a library for building API proxies. It can be used for API monitoring, Data anonymization and for multi format APIs.

Minimal Runnable example:   
(The code below will print new connection on a request)
```go
package main

import (
	"github.com/flew-software/Mesh-Layer/core"
	"net/http"
)

func main() {
    // create a mesh
	m1 := core.Mesh {
		Function: func(a core.MeshInput) {
			println("new connection")
			*a.Strand.Killed = false
		},
		Info: core.MeshInfo{},
	}

    // create a fabric(api endpoint)
	f := core.Fabric{
		Meshes: core.MeshRegistry{m1},
		Path:   "/test",
		RouteDest: core.API{
			Route: "/",
			Ip:    &core.Ip{Address: "https://api.mocki.io/v1/18c74236"},
			Code:  "",
		},
	}

    // create a cloth (you can add multiple fabrics)
	fl := core.Cloth{f}

	http.Handle("/", fl.Handle())
	http.ListenAndServe(":8080", nil)
}

```

 