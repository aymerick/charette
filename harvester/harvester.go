package harvester

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/aymerick/charette/rom"
)

var (
	// allowed file extensions
	extensions []string
)

func init() {
	// no-intro file extensions
	extensions = []string{".zip", ".7z"}
}

type Harvester struct {
	// roms directory
	Dir string

	Options *Options
	Debug   bool

	// found games
	Games map[string]*rom.Game
}

func New(dir string, options *Options) *Harvester {
	return &Harvester{
		Dir:     dir,
		Options: options,
		Debug:   options.Debug,
		Games:   map[string]*rom.Game{},
	}
}

func (h *Harvester) Run() error {
	if h.Debug {
		log.Printf("Scanning files: %s", h.Dir)
	}

	infos, err := ioutil.ReadDir(h.Dir)
	if err != nil {
		return err
	}

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
	}

	// @todo Filter roms

	return nil
}

func (h *Harvester) processFile(info os.FileInfo) error {
	if h.Debug {
		log.Printf("Processing: %s", info.Name())
	}

	r := rom.New(info.Name())
	if err := r.Fill(); err != nil {
		return err
	}

	if skip, msg := h.skip(r); skip {
		fmt.Printf("Skipped '%s': %s\n", r, msg)
		return nil
	}

	if h.Games[r.Name] == nil {
		// it's a new game
		g := rom.NewGame(r.Name)
		g.AddRom(r)

		h.Games[r.Name] = g

		if h.Debug {
			log.Printf("New game found: %s", g)
		}
	}

	if h.Debug {
		log.Printf("New rom found: %s", r)
	}

	return nil
}

func (h *Harvester) skip(r *rom.Rom) (bool, string) {
	if strings.HasPrefix(r.Name, "[BIOS]") {
		return true, "Ignore bios"
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
