package utils

import "strings"

// ExtractRegions returns an array of regions
func ExtractRegions(str string) []string {
	result := []string{}

	regions := strings.Split(str, ",")
	for _, region := range regions {
		result = append(result, strings.TrimSpace(region))
	}

	return result
}
