package rom

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/aymerick/charette/utils"
)

const (
	BIOS_PREFIX = "[BIOS]"
)

var rFilename = regexp.MustCompile(`^([^\(]*)\(([^\(]*)\)`)
var rProto = regexp.MustCompile(`\(Proto\)`)
var rBeta = regexp.MustCompile(`\(Beta\)`)

// Rom represents a game version
type Rom struct {
	Filename string
	Name     string
	Regions  []string
	Revision int

	Proto       bool
	Beta        bool
	BetaVersion int
	Bios        bool
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

// Fill extracts Rom infos from filename
func (r *Rom) Fill() error {
	// extract infos from filename
	match := rFilename.FindStringSubmatch(r.Filename)
	if len(match) != 3 {
		return errors.New("Invalid filename: " + r.Filename)
	}

	r.Name = strings.TrimSpace(match[1])
	r.Regions = utils.ExtractRegions(match[2])
	// @todo r.Revision

	r.Proto = rProto.MatchString(r.Filename)
	r.Beta = rBeta.MatchString(r.Filename)
	// @todo r.BetaVersion
	r.Bios = strings.HasPrefix(r.Filename, BIOS_PREFIX)

	return nil
}
