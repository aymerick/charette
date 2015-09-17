package rom

import (
	"fmt"
	"testing"
)

func TestGameBestRom(t *testing.T) {
	g := NewGame()
	g.AddRom(MustFill("Addams Family, The - Pugsley's Scavenger Hunt (USA).zip"))
	g.AddRom(MustFill("Addams Family, The - Pugsley's Scavenger Hunt (Europe) (Rev 1).zip"))
	g.AddRom(MustFill("Addams Family, The - Pugsley's Scavenger Hunt (Europe).zip"))
	r4 := g.AddRom(MustFill("Addams Family, The - Pugsley's Scavenger Hunt (Europe) (Rev 2).zip"))
	g.AddRom(MustFill("Addams Family, The - Pugsley's Scavenger Hunt (USA) (Beta).zip"))

	regions := []string{"Europe", "USA", "Japan"}

	if r := g.BestRom(regions); r != r4 {
		t.Errorf("Game best rom computation failed, got '%v' but expected '%v'", r, r4)
	}
}

func TestGameRomsSort(t *testing.T) {
	g := NewGame()

	r1 := g.AddRom(MustFill("Addams Family, The - Pugsley's Scavenger Hunt (Japan).zip"))
	r2 := g.AddRom(MustFill("Addams Family, The - Pugsley's Scavenger Hunt (USA).zip"))
	r3 := g.AddRom(MustFill("Addams Family, The - Pugsley's Scavenger Hunt (Europe) (Rev 1).zip"))
	r4 := g.AddRom(MustFill("Addams Family, The - Pugsley's Scavenger Hunt (Europe).zip"))
	r5 := g.AddRom(MustFill("Addams Family, The - Pugsley's Scavenger Hunt (Europe) (Rev 2).zip"))
	r6 := g.AddRom(MustFill("Addams Family, The - Pugsley's Scavenger Hunt (USA) (Beta 2).zip"))
	r7 := g.AddRom(MustFill("Addams Family, The - Pugsley's Scavenger Hunt (Japan) (Demo).zip"))
	r8 := g.AddRom(MustFill("Addams Family, The - Pugsley's Scavenger Hunt (USA) (Beta 1).zip"))

	regions := []string{"Europe", "USA", "Japan"}
	g.sortRoms(regions)

	expected := []*Rom{r5, r3, r4, r2, r6, r8, r1, r7}

	for i, rom := range g.Roms {
		if rom != expected[i] {
			t.Fatal(fmt.Sprintf("Game roms sort failed\n\tgot     : %v\n\texpected: %v", g.Roms, expected))
		}
	}
}

func TestGameRomsSort2(t *testing.T) {
	g := NewGame()

	r1 := g.AddRom(MustFill("Donkey Kong Country 2 - Diddy's Kong Quest (Europe) (Rev 1).zip"))
	r2 := g.AddRom(MustFill("Donkey Kong Country 2 - Diddy's Kong Quest (Germany).zip"))
	r3 := g.AddRom(MustFill("Donkey Kong Country 2 - Diddy's Kong Quest (Germany) (Rev 1).zip"))
	r4 := g.AddRom(MustFill("Donkey Kong Country 2 - Diddy's Kong Quest (USA).zip"))
	r5 := g.AddRom(MustFill("Donkey Kong Country 2 - Diddy's Kong Quest (USA) (Rev 1).zip"))

	regions := []string{"France", "Europe", "World", "USA", "Japan"}
	g.sortRoms(regions)

	expected := []*Rom{r1, r5, r4, r3, r2}

	for i, rom := range g.Roms {
		if rom != expected[i] {
			t.Fatal(fmt.Sprintf("Game roms sort failed\n\tgot     : %v\n\texpected: %v", g.Roms, expected))
		}
	}
}
