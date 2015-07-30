package rom

type Game struct {
	Name string
	Roms []*Rom
}

func NewGame(name string) *Game {
	return &Game{
		Name: name,
	}
}

func (g *Game) String() string {
	return g.Name
}

func (g *Game) AddRom(r *Rom) {
	g.Roms = append(g.Roms, r)
}
