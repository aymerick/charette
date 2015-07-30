package rom

// Games represents a game with multiple versions
type Game struct {
	Name string
	Roms []*Rom
}

// NewGame instanciates a new Game
func NewGame(name string) *Game {
	return &Game{
		Name: name,
	}
}

// String returns the string representation of a Game
func (g *Game) String() string {
	return g.Name
}

// AddRom adds a new game version
func (g *Game) AddRom(r *Rom) {
	g.Roms = append(g.Roms, r)
}
