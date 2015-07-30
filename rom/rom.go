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

// regexps
var rFilename = regexp.MustCompile(`^([^\(]*)\(([^\(]*)\)`)
var rVersion = regexp.MustCompile(`\(Rev(.*)\)|\(v\d.*\)`)

var rProto = regexp.MustCompile(`\(Proto\)`)
var rBeta = regexp.MustCompile(`\(Beta([^\(]*)\)`)
var rSample = regexp.MustCompile(`\(Sample\)`)
var rDemo = regexp.MustCompile(`\(([^\(]*)Demo([^\(]*)\)`)
var rPirate = regexp.MustCompile(`\(([^\(]*)Pirate([^\(]*)\)`)
var rPromo = regexp.MustCompile(`\(([^\(]*)Promo([^\(]*)\)`)

// Rom represents a game version
type Rom struct {
	Filename string
	Name     string
	Regions  []string
	Version  string

	Proto  bool
	Beta   bool
	Bios   bool
	Sample bool
	Demo   bool
	Pirate bool
	Promo  bool
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
	r.Version = r.extractVersion()

	r.Proto = rProto.MatchString(r.Filename)
	r.Beta = rBeta.MatchString(r.Filename)
	r.Bios = strings.HasPrefix(r.Filename, BIOS_PREFIX)
	r.Sample = rSample.MatchString(r.Filename)
	r.Demo = rDemo.MatchString(r.Filename)
	r.Pirate = rPirate.MatchString(r.Filename)
	r.Promo = rPromo.MatchString(r.Filename)

	return nil
}

func (r *Rom) extractVersion() string {
	result := ""

	match := rVersion.FindStringSubmatch(r.Filename)
	if len(match) == 2 {
		result = match[0]
	} else {
		match = rBeta.FindStringSubmatch(r.Filename)
		if len(match) == 2 {
			result = match[0]
		}
	}

	if result != "" {
		// drop parenthesis
		result = result[1 : len(result)-1]
	}

	return result
}
