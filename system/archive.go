package system

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/aymerick/charette/helpers"
	"github.com/aymerick/charette/rom"
)

// Archive represents an archive of system roms
type Archive struct {
	// gaming system
	System *System

	// archive path
	Path string

	// output directory path
	Output string

	// selected games
	Games map[string]*rom.Game

	// processed files number
	Processed int

	// skipped files number
	Skipped int

	// selected regions stats
	RegionsStats map[string]int
}

// NewArchive instanciates a new Archive
func NewArchive(s *System, path string, output string) *Archive {
	return &Archive{
		System:       s,
		Path:         path,
		Output:       output,
		Games:        map[string]*rom.Game{},
		RegionsStats: map[string]int{},
	}
}

// Process filters roms in archive
func (a *Archive) Process() error {
	workingDir := path.Join(a.System.Options.Tmp, helpers.FileBase(a.Path))

	args := []string{"x", a.Path, "-o" + workingDir, "-y"}
	if err := helpers.ExecCmd("7z", args); err != nil {
		return err
	}

	// process roms
	if err := a.SelectRoms(workingDir); err != nil {
		return err
	}

	// delete extracted archive directory
	if a.System.Options.Debug {
		fmt.Printf("[%s] Deleting temp dir: %s\n", a.System.Infos.Name, workingDir)
	}

	if err := os.RemoveAll(workingDir); err != nil {
		return err
	}

	return nil
}

// SelectRoms filters all roms found in input directory and copy selected ones to output directory
func (a *Archive) SelectRoms(inputDir string) error {
	// process input files
	if err := a.processDir(inputDir); err != nil {
		return err
	}

	// move selected files to output directory
	if err := a.moveSelectedRoms(); err != nil {
		return err
	}

	return nil
}

func (a *Archive) processDir(inputDir string) error {
	files, err := ioutil.ReadDir(inputDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			// @todo Ignore /roms and /.~charette directories !

			subdir := path.Join(inputDir, file.Name())

			// process subdirectory
			if err := a.processDir(subdir); err != nil {
				fmt.Printf("ERR: %v\n", err)
			}
		} else {
			// process file
			if err := a.processFile(inputDir, file); err != nil {
				fmt.Printf("ERR: %v\n", err)
			}

			a.Processed++
		}
	}

	return nil
}

func (a *Archive) processFile(inputDir string, file os.FileInfo) error {
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

	if skip, msg := a.skip(r); skip {
		if a.System.Options.Debug {
			fmt.Printf("[%s] Skipped '%s': %s\n", a.System.Infos.Name, r.Filename, msg)
		}

		a.Skipped++

		return nil
	}

	if a.Games[r.Name] == nil {
		// it's a new game
		a.Games[r.Name] = rom.NewGame()
	}

	a.Games[r.Name].AddRom(r)

	return nil
}

// skip returns true if given rom must be skiped, with an explanation message
func (a *Archive) skip(r *rom.Rom) (bool, string) {
	if a.System.Options.LeaveMeAlone && !r.HaveRegion(a.System.Options.Regions) {
		return true, fmt.Sprintf("Leave me alone: %v\n", r.Regions)
	}

	if r.Proto && !a.System.Options.KeepProto {
		return true, "Ignore proto"
	}

	if r.Beta && !a.System.Options.KeepBeta {
		return true, "Ignore beta"
	}

	if r.Bios {
		return true, "Ignore bios"
	}

	if r.Sample && !a.System.Options.KeepSample {
		return true, "Ignore sample"
	}

	if r.Demo && !a.System.Options.KeepDemo {
		return true, "Ignore demo"
	}

	if r.Pirate && !a.System.Options.KeepPirate {
		return true, "Ignore pirate"
	}

	if r.Promo && !a.System.Options.KeepPromo {
		return true, "Ignore promo"
	}

	return false, ""
}

// move selected roms to output directory
func (a *Archive) moveSelectedRoms() error {
	if a.System.Options.Debug {
		fmt.Printf("[%s] Moving all %v selected roms to: %s\n", a.System.Infos.Name, len(a.Games), a.Output)
	}

	for _, g := range a.Games {
		g.SortRoms(a.System.Options.Regions)
		r := g.BestRom()

		outputPath := path.Join(a.Output, r.Filename)

		if a.System.Options.Debug {
			fmt.Printf("MOVING: %s => %s\n", r.File, outputPath)
		}

		if err := os.Rename(r.File, outputPath); err != nil {
			return err
		}

		a.RegionsStats[r.BestRegion(a.System.Options.Regions)]++
	}

	return nil
}
