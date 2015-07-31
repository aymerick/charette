package rom

import "sort"

// Games represents a game with multiple versions
type Game struct {
	Name string
	Roms []*Rom

	sorted bool
}

// NewGame instanciates a new Game
func NewGame() *Game {
	return &Game{}
}

// String returns the string representation of a Game
func (g *Game) String() string {
	return g.Name
}

// AddRom adds a new game version
func (g *Game) AddRom(r *Rom) *Rom {
	if g.Name == "" {
		g.Name = r.Name
	}

	g.Roms = append(g.Roms, r)

	return r
}

// SortRoms sorts roms
func (g *Game) SortRoms(regions []string) {
	sort.Sort(g.NewRomsSort(regions))

	g.sorted = true
}

// BestRoom returns the best rom matching the given criteria
func (g *Game) BestRom() *Rom {
	if !g.sorted {
		panic("Roms must be sorted")
	}

	return g.Roms[0]
}

// BestRoom returns roms that can be discarded
func (g *Game) GarbageRoms() []*Rom {
	if !g.sorted {
		panic("Roms must be sorted")
	}

	if len(g.Roms) > 1 {
		return g.Roms[1:len(g.Roms)]
	}

	return []*Rom{}
}

//
// Sort
//

// Games represents a game with multiple versions
type GameRomsSort struct {
	Game    *Game
	Regions []string
}

func (g *Game) NewRomsSort(regions []string) *GameRomsSort {
	return &GameRomsSort{
		Game:    g,
		Regions: regions,
	}
}

// Implements sort.Interface
func (gs GameRomsSort) Len() int {
	return len(gs.Game.Roms)
}

// Implements sort.Interface
func (gs GameRomsSort) Swap(i, j int) {
	gs.Game.Roms[i], gs.Game.Roms[j] = gs.Game.Roms[j], gs.Game.Roms[i]
}

// Implements sort.Interface
func (gs GameRomsSort) Less(i, j int) bool {
	r1 := gs.Game.Roms[i]
	r2 := gs.Game.Roms[j]

	b1 := r1.BestRegionIndex(gs.Regions)
	b2 := r2.BestRegionIndex(gs.Regions)

	if b1 != b2 {
		return b1 < b2
	}

	// tag - any alternative tag is a looser
	if r1.HaveAltTag() != r2.HaveAltTag() {
		return r2.HaveAltTag()
	}

	// version - latest version is the winner
	if r1.Version != r2.Version {
		return r1.Version > r2.Version
	}

	return false
}
