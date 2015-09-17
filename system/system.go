package system

import (
	"os"
	"path"

	"github.com/aymerick/charette/core"
	"github.com/aymerick/charette/rom"
)

// System represents a gaming system found in no-intro archives
type System struct {
	// system informations
	Infos Infos

	// options
	Options *core.Options

	// all selected games from all archives
	Games map[string]*rom.Game

	// total number of processed files from all archives
	Processed int

	// total number of skipped files from all archives
	Skipped int

	// selected regions stats from all archives
	RegionsStats map[string]int
}

// New instanciates a new System
func New(infos Infos, options *core.Options) *System {
	return &System{
		Infos:        infos,
		Options:      options,
		Games:        map[string]*rom.Game{},
		RegionsStats: map[string]int{},
	}
}

// RomsDir returns the roms directory name for that system
func (s *System) RomsDir() string {
	return s.Infos.Dir
}

// ProcessArchive filters roms in given no-intro archive and outputs selected ones into given output directory
func (s *System) ProcessArchive(archive string, outputDir string) error {
	// ensure output directory
	outputDir = path.Join(outputDir, s.RomsDir())
	if err := os.MkdirAll(outputDir, 0777); (err != nil) && (err != os.ErrExist) {
		return err
	}

	// process archive
	a := NewArchive(s, archive, outputDir, s.Options)
	if err := a.Process(); err != nil {
		return err
	}

	// merge results
	for name, game := range a.Games {
		s.Games[name] = game
	}

	s.Processed += a.Processed
	s.Skipped += a.Skipped

	for region, nb := range a.RegionsStats {
		s.RegionsStats[region] += nb
	}

	return nil
}
