package rom

import (
	"fmt"
	"log"
	"testing"
)

func TestGameBestRom(t *testing.T) {
	g := NewGame()
	r1 := g.AddRom(MustFill("Addams Family, The - Pugsley's Scavenger Hunt (USA).zip"))
	r2 := g.AddRom(MustFill("Addams Family, The - Pugsley's Scavenger Hunt (Europe) (Rev 1).zip"))
	r3 := g.AddRom(MustFill("Addams Family, The - Pugsley's Scavenger Hunt (Europe).zip"))
	r4 := g.AddRom(MustFill("Addams Family, The - Pugsley's Scavenger Hunt (Europe) (Rev 2).zip"))
	r5 := g.AddRom(MustFill("Addams Family, The - Pugsley's Scavenger Hunt (USA) (Beta).zip"))

	expectedGarbage := []*Rom{r2, r3, r1, r5}

	regions := []string{"Europe", "USA", "Japan"}

	if r := g.BestRom(regions); r != r4 {
		t.Errorf("Game best rom computation failed, got '%v' but expected '%v'", r, r4)
	}

	garbage := g.GarbageRoms(regions)
	if len(garbage) != len(expectedGarbage) {
		t.Fatal(fmt.Sprintf("Game garbage rom computation failed\n\tgot     : %v\n\texpected: %v", garbage, expectedGarbage))
	}

	for i, rom := range garbage {
		if rom != expectedGarbage[i] {
			t.Fatal(fmt.Sprintf("Game garbage rom computation failed\n\tgot     : %v\n\texpected: %v", garbage, expectedGarbage))
		}
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

func TestGameRomsSort2(t *testing.T) {
	g := NewGame()

	r1 := g.AddRom(MustFill("Donkey Kong Country 2 - Diddy's Kong Quest (Europe) (Rev 1).zip"))
	r2 := g.AddRom(MustFill("Donkey Kong Country 2 - Diddy's Kong Quest (Germany).zip"))
	r3 := g.AddRom(MustFill("Donkey Kong Country 2 - Diddy's Kong Quest (Germany) (Rev 1).zip"))
	r4 := g.AddRom(MustFill("Donkey Kong Country 2 - Diddy's Kong Quest (USA).zip"))
	r5 := g.AddRom(MustFill("Donkey Kong Country 2 - Diddy's Kong Quest (USA) (Rev 1).zip"))

	regions := []string{"France", "Europe", "World", "USA", "Japan"}
	expected := []*Rom{r1, r5, r4, r3, r2}

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
