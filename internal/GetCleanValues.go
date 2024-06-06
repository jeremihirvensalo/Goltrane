package internal

import (
	"regexp"
	"strings"
)

var CleanStrings = map[string]string{
	"rgb":  `(\d{1,3} *, *\d{1,3} *, *\d{1,3})`,
	"rgba": `(\d{1,3} *, *\d{1,3} *, *\d{1,3} *, *[0-1]?\.{0,1}\d{1,2}?)`,
	"hsl":  `(-?\d{1,3} *, *\d{1,3}%? *, *\d{1,3}%)`,
}

func GetCleanValues(colorType string, hits []string) []string {
	reg := regexp.MustCompile(CleanStrings[colorType])

	cleanVals := make([]string, len(hits))

	for i := 0; i < len(hits); i++ {
		match := reg.FindString(hits[i])

		cleanVals[i] = strings.Replace(match, " ", "", -1)
	}

	return cleanVals
}
