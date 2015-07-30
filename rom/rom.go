package rom

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/aymerick/charette/utils"
)

var rFilename = regexp.MustCompile(`^([^\(]*)\(([^\(]*)\)`)

type Rom struct {
	Filename string
	Name     string
	Regions  []string
}

func New(filename string) *Rom {
	return &Rom{
		Filename: filename,
	}
}

func (r *Rom) String() string {
	return fmt.Sprintf("%s %v", r.Name, r.Regions)
}

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
