package core

import "strings"

var (
	// allowed regions
	AllowedRegions map[string]bool
)

func init() {
	AllowedRegions = map[string]bool{
		"Asia":        true,
		"Australia":   true,
		"Brazil":      true,
		"Canada":      true,
		"China":       true,
		"Denmark":     true,
		"Europe":      true,
		"Finland":     true,
		"France":      true,
		"Germany":     true,
		"Hong Kong":   true,
		"Italy":       true,
		"Japan":       true,
		"Korea":       true,
		"Netherlands": true,
		"Russia":      true,
		"Spain":       true,
		"Sweden":      true,
		"Taiwan":      true,
		"Unknown":     true,
		"USA":         true,
		"World":       true,
	}
}

// ExtractRegions returns an array of regions
func ExtractRegions(str string) []string {
	result := []string{}

	regions := strings.Split(str, ",")
	for _, region := range regions {
		region = strings.TrimSpace(region)

		if AllowedRegions[region] {
			result = append(result, region)
		}
	}

	return result
}
