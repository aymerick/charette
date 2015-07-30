package rom

import "testing"

var fillTests = []struct {
	fileName    string
	name        string
	regions     []string
	revision    int
	proto       bool
	beta        bool
	betaVersion int
	bios        bool
}{
	{
		"[BIOS] CX4 (World).zip",
		"[BIOS] CX4",
		[]string{"World"},
		0,
		false,
		false,
		0,
		true,
	},
	{
		"Bubsy in Claws Encounters of the Furred Kind (USA) (Beta 1).zip",
		"Bubsy in Claws Encounters of the Furred Kind",
		[]string{"USA"},
		0,
		false,
		true,
		1,
		false,
	},
	{
		"BS-X - Sore wa Namae o Nusumareta Machi no Monogatari (Japan) (Rev 1)",
		"BS-X - Sore wa Namae o Nusumareta Machi no Monogatari",
		[]string{"Japan"},
		1,
		false,
		false,
		0,
		false,
	},
	{
		"Capcom's Soccer Shootout (USA) (Beta)",
		"Capcom's Soccer Shootout",
		[]string{"USA"},
		0,
		false,
		true,
		0,
		false,
	},
	{
		"Captain Novolin (USA) (En,Fr,Es)",
		"Captain Novolin",
		[]string{"USA"},
		0,
		false,
		false,
		0,
		false,
	},
}

func TestFill(t *testing.T) {
	for _, test := range fillTests {
		rom := New(test.fileName)
		if err := rom.Fill(); err != nil {
			t.Fatal("Fill failed: %s", err)
		}

		if rom.Name != test.name {
			t.Errorf("Name extraction failed, got '%v' but expected '%v': %s", rom.Name, test.name, test.fileName)
		}

		if !testEq(rom.Regions, test.regions) {
			t.Errorf("Regions extraction failed, got '%v' but expected '%v': %s", rom.Regions, test.regions, test.fileName)
		}

		if rom.Revision != test.revision {
			t.Errorf("Revision extraction failed, got '%v' but expected '%v': %s", rom.Revision, test.revision, test.fileName)
		}

		if rom.Proto != test.proto {
			t.Errorf("Proto extraction failed, got '%v' but expected '%v': %s", rom.Proto, test.proto, test.fileName)
		}

		if rom.Beta != test.beta {
			t.Errorf("Beta extraction failed, got '%v' but expected '%v': %s", rom.Beta, test.beta, test.fileName)
		}

		if rom.BetaVersion != test.betaVersion {
			t.Errorf("Beta Version extraction failed, got '%v' but expected '%v': %s", rom.BetaVersion, test.betaVersion, test.fileName)
		}

		if rom.Bios != test.bios {
			t.Errorf("Bios extraction failed, got '%v' but expected '%v': %s", rom.Bios, test.bios, test.fileName)
		}
	}
}

func testEq(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
