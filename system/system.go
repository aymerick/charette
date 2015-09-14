package system

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/aymerick/charette/core"
	"github.com/aymerick/charette/rom"
)

// System represents a gaming system found in no-intro archives
type System struct {
	// system informations
	Infos Infos

	// options
	Options *core.Options

	// found games
	Games map[string]*rom.Game

	// processed files number
	Processed int

	// skipped files number
	Skipped int

	// selected regions stats
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

// SelectRoms filters all roms found in input directory and copy selected ones to output directory
func (s *System) SelectRoms(inputDir string, outputDir string) error {
	if s.Options.Verbose {
		fmt.Printf("[%s] Selecting roms\n", s.Infos.Name)
	}

	// process input files
	if err := s.processDir(inputDir, outputDir); err != nil {
		return err
	}

	if s.Options.Verbose {
		fmt.Printf("[%s] Processed %v files (skipped: %v)\n", s.Infos.Name, s.Processed, s.Skipped)
	}

	fmt.Printf("[%s] Selected %v games\n", s.Infos.Name, len(s.Games))

	// move selected files to output directory
	if err := s.moveSelectedRoms(outputDir); err != nil {
		return err
	}

	return nil
}

func (s *System) processDir(inputDir string, outputDir string) error {
	files, err := ioutil.ReadDir(inputDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			// @todo Ignore /roms and /.~charette directories !

			subdir := path.Join(inputDir, file.Name())

			// process subdirectory
			if err := s.processDir(subdir, outputDir); err != nil {
				fmt.Printf("ERR: %v\n", err)
			}
		} else {
			// process file
			if err := s.processFile(inputDir, file, outputDir); err != nil {
				fmt.Printf("ERR: %v\n", err)
			}

			s.Processed++
		}
	}

	return nil
}

func (s *System) processFile(inputDir string, file os.FileInfo, outputDir string) error {
	// check file type
	fileExt := filepath.Ext(file.Name())
	if (fileExt != ".zip") && (fileExt != ".7z") {
		// skip file
		return nil
	}

	filePath := path.Join(inputDir, file.Name())
	r := rom.New(filePath)
	if err := r.Fill(); err != nil {
		return err
	}

	if skip, msg := s.skip(r); skip {
		if s.Options.Debug {
			fmt.Printf("[%s] Skipped '%s': %s\n", s.Infos.Name, r.Filename, msg)
		}

		s.Skipped++

		return nil
	}

	if s.Games[r.Name] == nil {
		// it's a new game
		s.Games[r.Name] = rom.NewGame()
	}

	s.Games[r.Name].AddRom(r)

	return nil
}

// skip returns true if given rom must be skiped, with an explanation message
func (s *System) skip(r *rom.Rom) (bool, string) {
	if s.Options.LeaveMeAlone && !r.HaveRegion(s.Options.Regions) {
		return true, fmt.Sprintf("Leave me alone: %v\n", r.Regions)
	}

	if r.Proto && !s.Options.KeepProto {
		return true, "Ignore proto"
	}

	if r.Beta && !s.Options.KeepBeta {
		return true, "Ignore beta"
	}

	if r.Bios {
		return true, "Ignore bios"
	}

	if r.Sample && !s.Options.KeepSample {
		return true, "Ignore sample"
	}

	if r.Demo && !s.Options.KeepDemo {
		return true, "Ignore demo"
	}

	if r.Pirate && !s.Options.KeepPirate {
		return true, "Ignore pirate"
	}

	if r.Promo && !s.Options.KeepPromo {
		return true, "Ignore promo"
	}

	return false, ""
}

// move selected roms to output directory
func (s *System) moveSelectedRoms(outputDir string) error {
	if s.Options.Verbose {
		fmt.Printf("[%s] Moving all %v selected roms to: %s\n", s.Infos.Name, len(s.Games), outputDir)
	}

	for _, g := range s.Games {
		g.SortRoms(s.Options.Regions)
		r := g.BestRom()

		outputPath := path.Join(outputDir, r.Filename)

		if s.Options.Debug {
			fmt.Printf("MOVING: %s => %s\n", r.File, outputPath)
		}

		if err := os.Rename(r.File, outputPath); err != nil {
			return err
		}

		s.RegionsStats[r.BestRegion(s.Options.Regions)]++
	}

	return nil
}
