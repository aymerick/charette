package harvester

import (
	"fmt"
	"io/ioutil"
	"os"
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
		fmt.Printf("Scaning input dir: %s\n", h.Input)
	}

	// detect all no-intro archives
	systems, err := h.scanArchives(h.Input)
	if err != nil {
		return err
	}

	// process archives
	for infos, archives := range systems {
		system := h.addSystem(infos)

		if err := h.processSystemArchives(system, archives); err != nil {
			return err
		}
	}

	// Display stats
	if h.Options.Verbose {
		processed := 0
		skipped := 0
		games := 0

		for _, system := range h.Systems {
			processed += system.Processed
			skipped += system.Skipped
			games += len(system.Games)
		}

		fmt.Printf("Processed %v files (skipped: %v)\n", processed, skipped)
		fmt.Printf("Selected %v games\n", games)
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
				fmt.Printf("Scaning subdir: %s\n", filePath)
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

// processSystemArchives processes archives for given system
func (h *Harvester) processSystemArchives(system *system.System, archives []string) error {
	workingDir := path.Join(h.Options.Tmp, system.Infos.Name)

	// ensure output directory
	outputDir := path.Join(h.Output, system.RomsDir())
	if err := os.MkdirAll(outputDir, 0777); (err != nil) && (err != os.ErrExist) {
		return err
	}

	// extract archives
	if h.Options.Verbose {
		fmt.Printf("[%s] Extracting %v archive(s)\n", system.Infos.Name, len(archives))
	}

	for _, archive := range archives {
		args := []string{"x", archive, "-o" + workingDir, "-y"}
		if err := helpers.ExecCmd("7z", args); err != nil {
			return err
		}
	}

	// process roms
	if err := system.SelectRoms(workingDir, outputDir); err != nil {
		return err
	}

	// delete extracted archive directory
	if h.Options.Debug {
		fmt.Printf("Deleting temp dir: %s\n", workingDir)
	}

	if err := os.RemoveAll(workingDir); err != nil {
		return err
	}

	return nil
}
