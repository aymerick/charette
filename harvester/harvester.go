package harvester

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/aymerick/charette/rom"
)

var (
	// allowed file extensions
	extensions []string
)

func init() {
	// no-intro files extensions
	extensions = []string{
		".zip", ".7z", // compressed
		".a26",        // atari2600
		".fds",        // famicomdisksystem
		".gb", ".gbc", // gameboy
		".gba",        // gameboy advance
		".gg",         // gamegear
		".ms",         // mastersystem
		".gen", ".md", // megadrive
		".rom", ".msx1", ".msx2", // msx
		".nes",                 // nes
		".n64",                 // nintendo64
		".pce",                 // pcengine
		".iso", ".img", ".bin", // playstation
		".32x",         // sega32x
		".cue",         // segacd
		".sg",          // sg1000
		".smc", ".sfc", // snes
		".vb", // virtual boy
	}
}

// Harvester collects wanted roms from given directory
type Harvester struct {
	// roms directory
	Dir string

	// garbage directory
	Garbage string

	Options *Options
	Debug   bool

	// found games
	Games map[string]*rom.Game

	// skipped files
	Skipped []string
}

// New instanciates a new Harvester
func New(dir string, garbage string, options *Options) *Harvester {
	return &Harvester{
		Dir:     dir,
		Garbage: garbage,
		Options: options,
		Debug:   options.Debug,
		Games:   map[string]*rom.Game{},
	}
}

// Run detects roms in directory and filters them
func (h *Harvester) Run() error {
	if h.Debug {
		log.Printf("Scanning files: %s", h.Dir)
	}

	// process files
	infos, err := ioutil.ReadDir(h.Dir)
	if err != nil {
		return err
	}

	nb := 0

	for _, info := range infos {
		if info.IsDir() {
			// Skip directories
			continue
		}

		if !expectedExt(info) {
			continue
		}

		// process file
		if err := h.processFile(info); err != nil {
			fmt.Printf("ERR: %v\n", err)
		}

		nb += 1
	}

	fmt.Printf("Processed %v files\n", nb)
	fmt.Printf("Skipped %v files\n", len(h.Skipped))
	fmt.Printf("Found %v games\n", len(h.Games))

	// discard roms
	if err := h.discard(); err != nil {
		return err
	}

	return nil
}

// discard moves skiped and unwanted roms to garbage
func (h *Harvester) discard() error {
	if _, err := os.Stat(h.Garbage); os.IsNotExist(err) {
		// create garbage directory
		if err := os.MkdirAll(h.Garbage, 0777); err != nil {
			return err
		}
	}

	nb := 0

	// discard skipped files
	for _, fileName := range h.Skipped {
		h.moveFile(fileName)
	}

	// discard unwanted roms
	for _, game := range h.Games {
		roms := game.GarbageRoms(h.Options.Regions)
		for _, r := range roms {
			if err := h.moveFile(r.Filename); err != nil {
				fmt.Printf("ERR: %v\n", err)
			} else {
				nb += 1
			}
		}
	}

	if nb == 0 {
		fmt.Printf("No file moved\n")
	} else {
		fmt.Printf("Moved %v files to %v\n", nb, h.Garbage)
	}

	return nil
}

func (h *Harvester) moveFile(fileName string) error {
	oldPath := path.Join(h.Dir, fileName)
	newPath := path.Join(h.Garbage, fileName)

	if h.Options.Debug {
		log.Printf("MOVING %s => %s", oldPath, newPath)
	}

	if h.Options.Noop {
		// NOOP
		return nil
	}

	return os.Rename(oldPath, newPath)
}

// processFile handles a new file
func (h *Harvester) processFile(info os.FileInfo) error {
	r := rom.New(info.Name())
	if err := r.Fill(); err != nil {
		return err
	}

	if skip, msg := h.skip(r); skip {
		if h.Options.Verbose {
			fmt.Printf("Skipped '%s': %s\n", r.Filename, msg)
		}

		h.Skipped = append(h.Skipped, r.Filename)

		return nil
	}

	if h.Games[r.Name] == nil {
		// it's a new game
		h.Games[r.Name] = rom.NewGame()
	}

	h.Games[r.Name].AddRom(r)

	return nil
}

// skip returns true if given rom must be skiped, with an explanation message
func (h *Harvester) skip(r *rom.Rom) (bool, string) {
	if r.Proto && h.Options.NoProto {
		return true, "Ignore proto"
	}

	if r.Beta && h.Options.NoBeta {
		return true, "Ignore beta"
	}

	if r.Bios {
		return true, "Ignore bios"
	}

	if r.Sample && h.Options.NoSample {
		return true, "Ignore sample"
	}

	if r.Demo && h.Options.NoDemo {
		return true, "Ignore demo"
	}

	if r.Pirate && h.Options.NoPirate {
		return true, "Ignore pirate"
	}

	if r.Promo && h.Options.NoPromo {
		return true, "Ignore promo"
	}

	return false, ""
}

// expectedExt returns true if this is an expected no-intro rom file extension
func expectedExt(info os.FileInfo) bool {
	fileExt := filepath.Ext(info.Name())

	for _, ext := range extensions {
		if fileExt == ext {
			return true
		}
	}

	return false
}
