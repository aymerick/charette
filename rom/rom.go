package rom

import (
	"fmt"
	"path"
	"regexp"
	"strings"

	"github.com/aymerick/charette/core"
)

const (
	biosPrefix = "[BIOS]"
)

// regexps
var rTags = regexp.MustCompile(`(\([^\(]*\))`)
var rVersion = regexp.MustCompile(`\(Rev(.*)\)|\(v\d.*\)`)

var rProto = regexp.MustCompile(`\(Proto\)`)
var rBeta = regexp.MustCompile(`\(Beta([^\(]*)\)`)
var rSample = regexp.MustCompile(`\(Sample\)`)
var rDemo = regexp.MustCompile(`\(([^\(]*)Demo([^\(]*)\)`)
var rPirate = regexp.MustCompile(`\(([^\(]*)Pirate([^\(]*)\)`)
var rPromo = regexp.MustCompile(`\(([^\(]*)Promo([^\(]*)\)`)

// Rom represents a game version
type Rom struct {
	File     string
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
func New(filePath string) *Rom {
	return &Rom{
		File:     filePath,
		Filename: path.Base(filePath),
	}
}

// MustFill instanciates an fill a new Rom. Panics on error.
func MustFill(filePath string) *Rom {
	r := New(filePath)
	if err := r.Fill(); err != nil {
		panic(err)
	}
	return r
}

// HaveAltTag returns true if rom have an alternative version tag
func (r *Rom) HaveAltTag() bool {
	return r.Proto || r.Beta || r.Sample || r.Demo || r.Pirate || r.Promo
}

// String returns the string representation of Rom
func (r *Rom) String() string {
	result := ""

	if r.Bios {
		result += "[BIOS] "
	}

	result += fmt.Sprintf("%s (%s)", r.Name, strings.Join(r.Regions, ", "))

	if r.Proto || r.Sample || r.Demo || r.Pirate || r.Promo {
		tags := []string{}

		if r.Proto {
			tags = append(tags, "Proto")
		}

		if r.Sample {
			tags = append(tags, "Sample")
		}

		if r.Demo {
			tags = append(tags, "Demo")
		}

		if r.Pirate {
			tags = append(tags, "Pirate")
		}

		if r.Promo {
			tags = append(tags, "Promo")
		}

		result += " (" + strings.Join(tags, ", ") + ")"
	}

	if r.Version != "" {
		result += " (" + r.Version + ")"
	}

	return result
}

// Fill extracts Rom infos from filename
func (r *Rom) Fill() error {
	// extract all tags
	tags := rTags.FindAllStringIndex(r.Filename, -1)

	// find regions tag
	regionsIndex := -1

	for _, tag := range tags {
		// check tag (don't forget to remove parenthesis)
		r.Regions = core.ExtractRegions(r.Filename[tag[0]+1 : tag[1]-1])
		if len(r.Regions) > 0 {
			regionsIndex = tag[0]
			break
		}
	}

	if regionsIndex < 0 {
		fmt.Printf("[ERROR] Rom without region: %s", r.Filename)
		// should not happen
		r.Name = r.Filename
	} else {
		// everything before region tag is the rom name
		r.Name = r.Filename[0 : regionsIndex-1]
	}

	r.Name = strings.TrimSpace(r.Name)

	r.Version = r.extractVersion()

	r.Proto = rProto.MatchString(r.Filename)
	r.Beta = rBeta.MatchString(r.Filename)
	r.Bios = strings.HasPrefix(r.Filename, biosPrefix)
	r.Sample = rSample.MatchString(r.Filename)
	r.Demo = rDemo.MatchString(r.Filename)
	r.Pirate = rPirate.MatchString(r.Filename)
	r.Promo = rPromo.MatchString(r.Filename)

	return nil
}

// HaveRegion returns true if rom matches with given regions
func (r *Rom) HaveRegion(regions []string) bool {
	for _, region := range regions {
		if indexOf(r.Regions, region) != -1 {
			return true
		}
	}

	return false
}

// BestRegionIndex computes the lowest index in given regions list for that rom
func (r *Rom) BestRegionIndex(regions []string) int {
	result := len(regions)

	for _, region := range r.Regions {
		if i := indexOf(regions, region); (i != -1) && (i < result) {
			result = i
		}
	}

	return result
}

// BestRegion returns the best region name in given regions list for that rom
func (r *Rom) BestRegion(regions []string) string {
	i := r.BestRegionIndex(regions)
	if i < len(regions) {
		return regions[i]
	}

	return r.Regions[0]
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

// @todo Move that to a 'utils' package
func indexOf(ar []string, value string) int {
	for i, v := range ar {
		if v == value {
			return i
		}
	}

	// not found
	return -1
}
