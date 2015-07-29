package rom

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var rFilename = regexp.MustCompile(`^([^\(]*)\(([^\(]*)\)`)

type Version struct {
	Filename string
	Name     string
	Regions  []string
}

func NewVersion(filename string) *Version {
	return &Version{
		Filename: filename,
	}
}

func (v *Version) String() string {
	return fmt.Sprintf("%s %v", v.Name, v.Regions)
}

func (v *Version) Fill() error {
	// extract infos from filename
	match := rFilename.FindStringSubmatch(v.Filename)
	if len(match) != 3 {
		return errors.New("Invalid filename: " + v.Filename)
	}

	v.Name = strings.TrimSpace(match[1])

	regions := strings.Split(match[2], ",")
	for _, region := range regions {
		v.Regions = append(v.Regions, strings.TrimSpace(region))
	}

	return nil
}
