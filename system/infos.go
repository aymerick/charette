package system

import (
	"path"
	"strings"
)

// Infos represents infos about a gaming system
type Infos struct {
	Manufacturer string
	Name         string
}

// SupportedSystems holds infos for all supported systems
var SupportedSystems []Infos

// SupportedSystems holds infos for all supported systems, indexed by "<Manufacturer> - <Name>"
var SupportedSystemsMap map[string]Infos

func init() {
	SupportedSystems = []Infos{
		{"Atari", "5200"},
		{"Atari", "7800"},
		{"Atari", "Jaguar"},
		{"Atari", "Lynx"},
		{"Atari", "ST"},
		{"Bandai", "WonderSwan"},
		{"Bandai", "WonderSwan Color"},
		{"Casio", "Loopy"},
		{"Casio", "PV-1000"},
		{"Coleco", "ColecoVision"},
		{"Commodore", "64"},
		{"Commodore", "64 (PP)"},
		{"Commodore", "64 (Tapes)"},
		{"Commodore", "Amiga"},
		{"Commodore", "Plus-4"},
		{"Commodore", "VIC-20"},
		{"Emerson", "Arcadia 2001"},
		{"Entex", "Adventure Vision"},
		{"Epoch", "Super Cassette Vision"},
		{"Fairchild", "Channel F"},
		{"Funtech", "Super Acan"},
		{"GamePark", "GP32"},
		{"GCE", "Vectrex"},
		{"Hartung", "Game Master"},
		{"LeapFrog", "Leapster Learning Game System"},
		{"Magnavox", "Odyssey2"},
		{"Microsoft", "MSX"},
		{"Microsoft", "MSX 2"},
		{"NEC", "PC Engine - TurboGrafx 16"},
		{"NEC", "Super Grafx"},
		{"Nintendo", "Famicom Disk System"},
		{"Nintendo", "Game Boy"},
		{"Nintendo", "Game Boy Advance"},
		{"Nintendo", "Game Boy Color"},
		{"Nintendo", "Nintendo 64"},
		{"Nintendo", "Nintendo Entertainment System"},
		{"Nintendo", "Pokemon Mini"},
		{"Nintendo", "Satellaview"},
		{"Nintendo", "Sufami Turbo"},
		{"Nintendo", "Super Nintendo Entertainment System"},
		{"Nintendo", "Virtual Boy"},
		{"Nokia", "N-Gage"},
		{"Philips", "Videopac+"},
		{"RCA", "Studio II"},
		{"Sega", "32X"},
		{"Sega", "Game Gear"},
		{"Sega", "Master System - Mark III"},
		{"Sega", "Mega Drive - Genesis"},
		{"Sega", "PICO"},
		{"Sega", "SG-1000"},
		{"Sinclair", "ZX Spectrum +3"},
		{"SNK", "Neo Geo Pocket"},
		{"SNK", "Neo Geo Pocket Color"},
		{"Tiger", "Game.com"},
		{"Tiger", "Gizmondo"},
		{"VTech", "CreatiVision"},
		{"VTech", "V.Smile"},
		{"Watara", "Supervision"},
	}

	SupportedSystemsMap = make(map[string]Infos)
	for _, infos := range SupportedSystems {
		SupportedSystemsMap[infos.Manufacturer+" - "+infos.Name] = infos
	}
}

// InfosForArchive returns system informations corresponding to archive name, the second value returned is `false` if system was not found
func InfosForArchive(filePath string) (Infos, bool) {
	var result Infos
	found := false

	arr := strings.Split(path.Base(filePath), " (")
	if len(arr) == 2 {
		result = SupportedSystemsMap[arr[0]]
		found = true
	}

	return result, found
}
