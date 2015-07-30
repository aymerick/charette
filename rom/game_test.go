package rom

import (
	"fmt"
	"log"
	"testing"
)

func TestGameBestRom(t *testing.T) {
	g := NewGame()
	g.AddRom(MustFill("Addams Family, The - Pugsley's Scavenger Hunt (USA).zip"))
	g.AddRom(MustFill("Addams Family, The - Pugsley's Scavenger Hunt (Europe) (Rev 1).zip"))
	g.AddRom(MustFill("Addams Family, The - Pugsley's Scavenger Hunt (Europe).zip"))
	rExp := g.AddRom(MustFill("Addams Family, The - Pugsley's Scavenger Hunt (Europe) (Rev 2).zip"))
	g.AddRom(MustFill("Addams Family, The - Pugsley's Scavenger Hunt (USA) (Beta).zip"))

	regions := []string{"Europe", "USA", "Japan"}

	if r := g.BestRom(regions); r != rExp {
		t.Errorf("Game best rom computation failed, got '%v' but expected '%v'", r, rExp)
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
	expected := []*Rom{r5, r3, r4, r2, r6, r8, r1, r7}

	g.sortRoms(regions)
	for _, rom := range g.Roms {
		log.Printf("    %v", rom)
	}

	for i, rom := range g.Roms {
		if rom != expected[i] {
			t.Fatal(fmt.Sprintf("Game roms sort failed\n\tgot     : %v\n\texpected: %v", g.Roms, expected))
		}
	}
}
