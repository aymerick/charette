package rom

import "testing"

var fillTests = []struct {
	fileName string
	name     string
	regions  []string
	version  string

	proto  bool
	beta   bool
	bios   bool
	sample bool
}{
	{
		"[BIOS] CX4 (World).zip",
		"[BIOS] CX4",
		[]string{"World"},
		"",
		false, false, true, false,
	},
	{
		"Bubsy in Claws Encounters of the Furred Kind (USA) (Beta 1).zip",
		"Bubsy in Claws Encounters of the Furred Kind",
		[]string{"USA"},
		"Beta 1",
		false, true, false, false,
	},
	{
		"Aretha II - Ariel no Fushigi na Tabi (Japan) (Beta 2).zip",
		"Aretha II - Ariel no Fushigi na Tabi",
		[]string{"Japan"},
		"Beta 2",
		false, true, false, false,
	},
	{
		"BS-X - Sore wa Namae o Nusumareta Machi no Monogatari (Japan) (Rev 1)",
		"BS-X - Sore wa Namae o Nusumareta Machi no Monogatari",
		[]string{"Japan"},
		"Rev 1",
		false, false, false, false,
	},
	{
		"Capcom's Soccer Shootout (USA) (Beta)",
		"Capcom's Soccer Shootout",
		[]string{"USA"},
		"Beta",
		false, true, false, false,
	},
	{
		"Captain Novolin (USA) (En,Fr,Es)",
		"Captain Novolin",
		[]string{"USA"},
		"",
		false, false, false, false,
	},
	{
		"Adventures of Dr. Franken, The (Europe) (En,Fr,De,Es,It,Nl,Sv)",
		"Adventures of Dr. Franken, The",
		[]string{"Europe"},
		"",
		false, false, false, false,
	},
	{
		"Axelay (USA) (Sample).zip",
		"Axelay",
		[]string{"USA"},
		"",
		false, false, false, true,
	},
	{
		"Gain Ground (World) (Rev A).zip",
		"Gain Ground",
		[]string{"World"},
		"Rev A",
		false, false, false, false,
	},
	{
		"Mortal Kombat (World) (v1.1).zip",
		"Mortal Kombat",
		[]string{"World"},
		"v1.1",
		false, false, false, false,
	},
}

// @todo
//
//   Micro Machines (USA, Europe) (Alt 1).zip
//   Micro Machines (USA, Europe) (MDMM ACD3).zip
//   Mike Ditka Power Football (USA, Europe) (Unl).zip
//   NBA Showdown '94 (USA) (Unl) (Pirate).zip
//

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

		if rom.Version != test.version {
			t.Errorf("Version extraction failed, got '%v' but expected '%v': %s", rom.Version, test.version, test.fileName)
		}

		if rom.Proto != test.proto {
			t.Errorf("Proto extraction failed, got '%v' but expected '%v': %s", rom.Proto, test.proto, test.fileName)
		}

		if rom.Beta != test.beta {
			t.Errorf("Beta extraction failed, got '%v' but expected '%v': %s", rom.Beta, test.beta, test.fileName)
		}

		if rom.Bios != test.bios {
			t.Errorf("Bios extraction failed, got '%v' but expected '%v': %s", rom.Bios, test.bios, test.fileName)
		}

		if rom.Sample != test.sample {
			t.Errorf("Sample extraction failed, got '%v' but expected '%v': %s", rom.Sample, test.sample, test.fileName)
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
