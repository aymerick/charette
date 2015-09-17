package harvester

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/cheggaaa/pb"

	"github.com/aymerick/charette/core"
	"github.com/aymerick/charette/system"
)

// Harvester collects wanted roms from given directory
type Harvester struct {
	// options
	Options *core.Options

	// systems found
	Systems []*system.System
}

// New instanciates a new Harvester
func New(options *core.Options) *Harvester {
	return &Harvester{
		Options: options,
	}
}

// Run detects systems archives in input directory and processes them
func (h *Harvester) Run() error {
	if h.Options.Debug {
		fmt.Printf("Scaning input dir: %s\n", h.Options.Input)
	}

	// detect all no-intro archives
	systems, err := h.scanArchives(h.Options.Input)
	if err != nil {
		return err
	}

	// process archives
	for infos, archives := range systems {
		s := h.addSystem(infos)

		if err := h.processSystemArchives(s, archives); err != nil {
			return err
		}
	}

	// Display stats
	h.printStats()

	return nil
}

func (h *Harvester) printStats() {
	processed := 0
	skipped := 0
	games := 0
	regions := map[string]int{}

	for _, s := range h.Systems {
		processed += s.Processed
		skipped += s.Skipped
		games += len(s.Games)

		for region, nb := range s.RegionsStats {
			regions[region] += nb
		}
	}

	fmt.Printf("=============== TOTAL ===============\n")
	fmt.Printf("Processed %v files (skipped: %v)\n", processed, skipped)
	fmt.Printf("Selected %v games\n", games)
	fmt.Printf("Regions:\n")

	for region, nb := range regions {
		fmt.Printf("\t%s: %d\n", region, nb)
	}
}

// scanArchives returns a map of {System Infos} => [Archives paths]
func (h *Harvester) scanArchives(input string) (map[system.Infos][]string, error) {
	result := make(map[system.Infos][]string)

	fileInfo, err := os.Stat(input)
	if err != nil {
		return result, err
	}

	if fileInfo.IsDir() {
		// scan archives directory
		return h.scanArchivesDir(input)
	}

	// scan archive file
	infos, found := h.scanArchiveFile(input)
	if found {
		result[infos] = []string{input}
	}

	return result, nil
}

func (h *Harvester) scanArchivesDir(dirPath string) (map[system.Infos][]string, error) {
	result := make(map[system.Infos][]string)

	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return result, err
	}

	for _, file := range files {
		filePath := path.Clean(path.Join(dirPath, file.Name()))

		if file.IsDir() {
			// ignore /roms and /.~charette directories
			if (filePath != path.Clean(h.Options.Output)) && (filePath != path.Clean(h.Options.Tmp)) {
				// scan subdir
				if h.Options.Debug {
					fmt.Printf("Scaning subdir: %s\n", filePath)
				}

				subArchives, err := h.scanArchivesDir(filePath)
				if err != nil {
					return result, nil
				}

				for infos, archives := range subArchives {
					result[infos] = append(result[infos], archives...)
				}
			}
		} else {
			if infos, found := h.scanArchiveFile(filePath); found {
				result[infos] = append(result[infos], filePath)
			}
		}
	}

	return result, nil
}

func (h *Harvester) scanArchiveFile(filePath string) (system.Infos, bool) {
	var result system.Infos
	found := false

	// scan archive
	fileExt := filepath.Ext(filePath)
	if fileExt == ".7z" {
		result, found = system.InfosForArchive(filePath)
	}

	return result, found
}

// addSystem registers a new system
func (h *Harvester) addSystem(infos system.Infos) *system.System {
	result := system.New(infos, h.Options)

	h.Systems = append(h.Systems, result)

	return result
}

// processSystemArchives processes archives for given system
func (h *Harvester) processSystemArchives(s *system.System, archives []string) error {
	var bar *pb.ProgressBar

	nb := len(archives)

	// extract archives
	if !s.Options.Quiet {
		fmt.Printf("[%s] Extracting %v archive(s)\n", s.Infos.Name, nb)

		if !s.Options.Debug {
			bar = pb.StartNew(nb)
			bar.ShowCounters = true
			bar.ShowPercent = false
			bar.ShowTimeLeft = true
			bar.SetMaxWidth(80)
		}
	}

	for _, archive := range archives {
		if !s.Options.Quiet && !s.Options.Debug {
			bar.Increment()
		}

		if err := s.ProcessArchive(archive, h.Options.Output); err != nil {
			return err
		}
	}

	if !s.Options.Quiet && !s.Options.Debug {
		bar.Finish()
		fmt.Printf("[%s] Processed %v files (skipped: %v)\n", s.Infos.Name, s.Processed, s.Skipped)
	}

	fmt.Printf("[%s] Selected %v games\n", s.Infos.Name, len(s.Games))

	return nil
}
