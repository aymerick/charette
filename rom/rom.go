package rom

type Rom struct {
	Name     string
	Versions []*Version
}

func New(name string) *Rom {
	return &Rom{
		Name: name,
	}
}

func (r *Rom) String() string {
	return r.Name
}

func (r *Rom) AddVersion(version *Version) {
	r.Versions = append(r.Versions, version)
}
