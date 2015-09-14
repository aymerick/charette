package harvester

import (
	"io/ioutil"
	"log"
	"path"
	"path/filepath"

	"github.com/aymerick/charette/core"
	"github.com/aymerick/charette/helpers"
	"github.com/aymerick/charette/system"
)

// Harvester collects wanted roms from given directory
type Harvester struct {
	// input directory
	Input string

	// output directory
	Output string

	// temporary working directory
	Tmp string

	// options
	Options *core.Options

	// systems found
	Systems []*system.System
}

// New instanciates a new Harvester
func New(input string, output string, tmp string, options *core.Options) *Harvester {
	return &Harvester{
		Input:   input,
		Output:  output,
		Tmp:     tmp,
		Options: options,
	}
}

// Run detects systems archives in input directory and processes them
func (h *Harvester) Run() error {
	if h.Options.Verbose {
		log.Printf("Scaning input dir: %s", h.Input)
	}

	// detect all no-intro archives
	systems, err := h.scanArchives(h.Input)
	if err != nil {
		return err
	}

	// process archives
	for infos, archives := range systems {
		system := h.addSystem(infos)

		if err := h.processArchives(system, archives); err != nil {
			return err
		}
	}

	return nil
}

// scanArchives returns a map of {System Infos} => [Archives paths]
func (h *Harvester) scanArchives(dir string) (map[system.Infos][]string, error) {
	result := make(map[system.Infos][]string)

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return result, err
	}

	for _, file := range files {
		filePath := path.Join(dir, file.Name())

		if file.IsDir() {
			if h.Options.Verbose {
				log.Printf("Scanning subdir: %s", filePath)
			}

			subArchives, err := h.scanArchives(filePath)
			if err != nil {
				return result, nil
			}

			for infos, archives := range subArchives {
				result[infos] = append(result[infos], archives...)
			}
		} else {
			fileExt := filepath.Ext(filePath)
			if fileExt == ".7z" {
				if infos, found := system.InfosForArchive(filePath); found {
					result[infos] = append(result[infos], filePath)
				}
			}
		}
	}

	return result, nil
}

// addSystem registers a new system
func (h *Harvester) addSystem(infos system.Infos) *system.System {
	result := system.New(infos, h.Options)

	h.Systems = append(h.Systems, result)

	return result
}

// processArchives processes archives for given system
func (h *Harvester) processArchives(system *system.System, archives []string) error {
	for _, archive := range archives {
		baseName := helpers.FileBase(archive)
		workingDir := path.Join(h.Options.Tmp, baseName)

		if h.Options.Verbose {
			log.Printf("Extracting archive into: %s", workingDir)
		}

		if !h.Options.Noop {
			// extract archive
			args := []string{"x", archive, "-o" + workingDir}
			if err := helpers.ExecCmd("7z", args); err != nil {
				return err
			}
		}

		// @todo Process roms
	}

	return nil
}
