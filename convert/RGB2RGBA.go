package convert

import (
	"karhuhelsinki/goltrane/internal"
	"strings"
)

func RGB2RGBA(data string) string {
	hits := internal.Regex("rgb", data)
	cleanVals := internal.GetCleanValues("rgb", hits)

	for i := 0; i < len(hits); i++ {
		cleanVals[i] = cleanVals[i] + ",1"

		color := "rgba(" + strings.Replace(cleanVals[i], ",", ", ", -1) + ")"

		data = strings.Replace(data, hits[i], color, -1)
	}

	return data
}
