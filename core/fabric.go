package core

// Fabric struct
type Fabric struct {
	Meshes    MeshRegistry
	Path      string
	RouteDest API
	uid       string
	MeshCount uint
}
