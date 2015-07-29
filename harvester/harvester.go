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

	// found roms
	Roms map[string]*rom.Rom
}

func New(dir string, options *Options) *Harvester {
	return &Harvester{
		Dir:     dir,
		Options: options,
		Debug:   options.Debug,
		Roms:    map[string]*rom.Rom{},
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

	// @todo Filter roms (options: extract)

	return nil
}

func (h *Harvester) processFile(info os.FileInfo) error {
	if h.Debug {
		log.Printf("Processing: %s", info.Name())
	}

	v := rom.NewVersion(info.Name())
	if err := v.Fill(); err != nil {
		return err
	}

	if skip, msg := h.skip(v); skip {
		fmt.Printf("Skipped '%s': %s\n", v, msg)
		return nil
	}

	if h.Roms[v.Name] == nil {
		// it's a new rom
		r := rom.New(v.Name)
		r.AddVersion(v)

		h.Roms[v.Name] = r

		if h.Debug {
			log.Printf("New rom found: %s", r)
		}
	}

	if h.Debug {
		log.Printf("New version found: %s", v)
	}

	return nil
}

func (h *Harvester) skip(v *rom.Version) (bool, string) {
	if strings.HasPrefix(v.Name, "[BIOS]") {
		return true, "Ignore bios file"
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
