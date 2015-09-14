package system

import (
	"path"
	"strings"
)

// Infos represents infos about a gaming system
type Infos struct {
	Manufacturer string
	Name         string
	Dir          string
}

// SupportedSystems holds infos for all supported systems
var SupportedSystems []Infos

// SupportedSystems holds infos for all supported systems, indexed by "<Manufacturer> - <Name>"
var SupportedSystemsMap map[string]Infos

func init() {
	SupportedSystems = []Infos{
		{"Atari", "5200", "atari2600"},
		{"Atari", "7800", "atari7800"},
		{"Atari", "Jaguar", "atarijaguar"},
		{"Atari", "Lynx", "lynx"},
		{"Atari", "ST", "atarist"},
		{"Bandai", "WonderSwan", "wswan"},
		{"Bandai", "WonderSwan Color", "wswan"},
		{"Casio", "Loopy", "loopy"},
		{"Casio", "PV-1000", "pv1000"},
		{"Coleco", "ColecoVision", "colecovision"},
		{"Commodore", "64", "c64"},
		{"Commodore", "64 (PP)", "c64"},
		{"Commodore", "64 (Tapes)", "c64"},
		{"Commodore", "Amiga", "amiga"},
		{"Commodore", "Plus-4", "plus4"},
		{"Commodore", "VIC-20", "vic20"},
		{"Emerson", "Arcadia 2001", "arcadia2001"},
		{"Entex", "Adventure Vision", "adventurevision"},
		{"Epoch", "Super Cassette Vision", "supercassettevision"},
		{"Fairchild", "Channel F", "channelf"},
		{"Funtech", "Super Acan", "superacan"},
		{"GamePark", "GP32", "gp32"},
		{"GCE", "Vectrex", "vectrex"},
		{"Hartung", "Game Master", "gamemaster"},
		{"LeapFrog", "Leapster Learning Game System", "llgs"},
		{"Magnavox", "Odyssey2", "odyssey2"},
		{"Microsoft", "MSX", "msx"},
		{"Microsoft", "MSX 2", "msx"},
		{"NEC", "PC Engine - TurboGrafx 16", "pcengine"},
		{"NEC", "Super Grafx", "pcengine"},
		{"Nintendo", "Famicom Disk System", "fds"},
		{"Nintendo", "Game Boy", "gb"},
		{"Nintendo", "Game Boy Advance", "gba"},
		{"Nintendo", "Game Boy Color", "gbc"},
		{"Nintendo", "Nintendo 64", "n64"},
		{"Nintendo", "Nintendo Entertainment System", "nes"},
		{"Nintendo", "Pokemon Mini", "pm"},
		{"Nintendo", "Satellaview", "satellaview"},
		{"Nintendo", "Sufami Turbo", "sufamiturbo"},
		{"Nintendo", "Super Nintendo Entertainment System", "snes"},
		{"Nintendo", "Virtual Boy", "virtualboy"},
		{"Nokia", "N-Gage", "ngage"},
		{"Philips", "Videopac+", "videopac"},
		{"RCA", "Studio II", "studio2"},
		{"Sega", "32X", "sega32x"},
		{"Sega", "Game Gear", "gamegear"},
		{"Sega", "Master System - Mark III", "mastersystem"},
		{"Sega", "Mega Drive - Genesis", "megadrive"},
		{"Sega", "PICO", "pico"},
		{"Sega", "SG-1000", "sg1000"},
		{"Sinclair", "ZX Spectrum +3", "zxspectrum"},
		{"SNK", "Neo Geo Pocket", "ngp"},
		{"SNK", "Neo Geo Pocket Color", "ngp"},
		{"Tiger", "Game.com", "gamecom"},
		{"Tiger", "Gizmondo", "gizmondo"},
		{"VTech", "CreatiVision", "creativision"},
		{"VTech", "V.Smile", "vsmile"},
		{"Watara", "Supervision", "supervision"},
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
