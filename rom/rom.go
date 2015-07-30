package rom

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/aymerick/charette/utils"
)

var rFilename = regexp.MustCompile(`^([^\(]*)\(([^\(]*)\)`)

// Rom represents a game version
type Rom struct {
	Filename string
	Name     string
	Regions  []string
}

// New instanciates a new Rom
func New(filename string) *Rom {
	return &Rom{
		Filename: filename,
	}
}

// String returns the string representation of Rom
func (r *Rom) String() string {
	return fmt.Sprintf("%s %v", r.Name, r.Regions)
}

// Fill extract Rom infos from filename
func (r *Rom) Fill() error {
	// extract infos from filename
	match := rFilename.FindStringSubmatch(r.Filename)
	if len(match) != 3 {
		return errors.New("Invalid filename: " + r.Filename)
	}

	r.Name = strings.TrimSpace(match[1])
	r.Regions = utils.ExtractRegions(match[2])

	return nil
}
