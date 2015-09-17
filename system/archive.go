package system

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/aymerick/charette/core"
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

	// options
	Options *core.Options

	// working directory path
	WorkingDir string

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
func NewArchive(s *System, filePath string, output string, options *core.Options) *Archive {
	result := &Archive{
		System:       s,
		Path:         filePath,
		Output:       output,
		Options:      options,
		Games:        map[string]*rom.Game{},
		RegionsStats: map[string]int{},
	}

	result.WorkingDir = path.Join(options.Tmp, helpers.FileBase(filePath))

	return result
}

func (a *Archive) log(msg string) {
	fmt.Printf("[%s] %s", a.System.Infos.Name, msg)
}

// Process filters roms in archive
func (a *Archive) Process() error {
	// extract archive
	if err := a.extract(); err != nil {
		return err
	}

	// process roms
	if err := a.selectRoms(); err != nil {
		return err
	}

	// delete extracted archive directory
	if err := a.cleanup(); err != nil {
		return err
	}

	return nil
}

// extractFile extracts given archive file into given output directory
func (a *Archive) extractFile(filePath string, output string) error {
	args := []string{"x", filePath, "-o" + output, "-y"}

	if a.Options.Debug {
		a.log(fmt.Sprintf("Extracting '%s' into: %s\n", path.Base(filePath), output))
	}

	return helpers.ExecCmd("7z", args)
}

// extract extracts the archive into working directory
func (a *Archive) extract() error {
	return a.extractFile(a.Path, a.WorkingDir)
}

// selectRoms filters all roms found in working directory and copy selected ones to output directory
func (a *Archive) selectRoms() error {
	// process input files
	if err := a.processDir(a.WorkingDir); err != nil {
		return err
	}

	// move selected files to output directory
	if err := a.moveSelectedRoms(); err != nil {
		return err
	}

	return nil
}

// deleteDir deletes given directory files
func (a *Archive) deleteDir(dir string) error {
	if a.Options.Debug {
		a.log(fmt.Sprintf("Deleting directory: %s\n", dir))
	}

	return os.RemoveAll(dir)
}

// cleanup deletes all extracted files
func (a *Archive) cleanup() error {
	return a.deleteDir(a.WorkingDir)
}

// processDir processes files in given directory
func (a *Archive) processDir(inputDir string) error {
	files, err := ioutil.ReadDir(inputDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			subdir := path.Join(inputDir, file.Name())

			// process subdirectory
			if err := a.processDir(subdir); err != nil {
				a.log(fmt.Sprintf("ERR: %v\n", err))
			}
		} else {
			// process file
			if err := a.processFile(inputDir, file); err != nil {
				a.log(fmt.Sprintf("ERR: %v\n", err))
			}
		}
	}

	return nil
}

// processFile processes file at given path
func (a *Archive) processFile(inputDir string, file os.FileInfo) error {
	a.Processed++

	// check file type
	fileExt := filepath.Ext(file.Name())
	if (fileExt != ".zip") && (fileExt != ".7z") {
		// skip file
		return nil
	}

	filePath := path.Join(inputDir, file.Name())

	if fileExt == ".7z" {
		// this is an archive of a specific game, with potentially several roms from different regions
		return a.processGameArchive(filePath)
	}

	return a.processGameFile(filePath)
}

// processGameFile processes game file at given path
func (a *Archive) processGameFile(filePath string) error {
	r := rom.New(filePath)
	if err := r.Fill(); err != nil {
		return err
	}

	if skip, msg := a.skip(r); skip {
		if a.Options.Debug {
			a.log(fmt.Sprintf("Skipped '%s': %s\n", r.Filename, msg))
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

// processGameArchive processes game archive at given path
func (a *Archive) processGameArchive(filePath string) error {
	gamesDir := path.Join(a.WorkingDir, helpers.FileBase(filePath))

	// extract game archive
	if err := a.extractFile(filePath, gamesDir); err != nil {
		return err
	}

	// select one rom from game archive
	if err := a.selectGameArchiveRom(gamesDir); err != nil {
		return err
	}

	// delete extracted game archive
	if err := a.deleteDir(gamesDir); err != nil {
		return err
	}

	return nil
}

// selectGameArchiveRom selects only one rom from given game archive directory, mark it as already selected then copy it to output
func (a *Archive) selectGameArchiveRom(dir string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	g := rom.NewGame()

	// process roms
	for _, file := range files {
		if file.IsDir() {
			a.log(fmt.Sprintf("ERR: Skipped unexpected directory in game archive: %s\n", file.Name()))
		} else {
			filePath := path.Join(dir, file.Name())

			r := rom.New(filePath)
			if err := r.Fill(); err != nil {
				return err
			}

			if skip, msg := a.skip(r); skip {
				if a.Options.Debug {
					a.log(fmt.Sprintf("Skipped '%s': %s\n", r.Filename, msg))
				}

				a.Skipped++
			} else {
				g.AddRom(r)
			}

			a.Processed++
		}
	}

	// select best rom
	if err := a.moveGameBestRom(g); err != nil {
		return err
	}

	gName, _ := rom.NameAndRegions(path.Base(dir))
	a.Games[gName] = g

	return nil
}

// skip returns true if given rom must be skiped, with an explanation message
func (a *Archive) skip(r *rom.Rom) (bool, string) {
	if a.Options.Strict && !r.HaveRegion(a.Options.Regions) {
		return true, fmt.Sprintf("Strict: %v", r.Regions)
	}

	if r.Proto && !a.Options.KeepProto {
		return true, "Ignore proto"
	}

	if r.Beta && !a.Options.KeepBeta {
		return true, "Ignore beta"
	}

	if r.Bios {
		return true, "Ignore bios"
	}

	if r.Sample && !a.Options.KeepSample {
		return true, "Ignore sample"
	}

	if r.Demo && !a.Options.KeepDemo {
		return true, "Ignore demo"
	}

	if r.Pirate && !a.Options.KeepPirate {
		return true, "Ignore pirate"
	}

	if r.Promo && !a.Options.KeepPromo {
		return true, "Ignore promo"
	}

	return false, ""
}

// moveFile moves given file into given directory
func (a *Archive) moveFile(filePath string, dir string) error {
	if a.Options.Debug {
		a.log(fmt.Sprintf("Moving '%s' into: %s\n", filePath, dir))
	}

	return os.Rename(filePath, dir)
}

// moveGameBestRom moves best rom of given game to output directory
func (a *Archive) moveGameBestRom(g *rom.Game) error {
	if g.Moved {
		// game was already moved
		return nil
	}

	r := g.BestRom(a.Options.Regions)
	if r == nil {
		// no rom matches filtering criteria
		return nil
	}

	outputPath := path.Join(a.Output, r.Filename)

	if err := a.moveFile(r.File, outputPath); err != nil {
		return err
	}

	g.Moved = true

	a.RegionsStats[r.BestRegion(a.Options.Regions)]++

	return nil
}

// moveSelectedRoms moves selected roms to output directory
func (a *Archive) moveSelectedRoms() error {
	if a.Options.Debug {
		a.log(fmt.Sprintf("Moving all %v selected roms to: %s\n", len(a.Games), a.Output))
	}

	for _, g := range a.Games {
		if err := a.moveGameBestRom(g); err != nil {
			return err
		}
	}

	return nil
}
