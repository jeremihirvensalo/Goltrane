package convert

import (
	"karhuhelsinki/goltrane/internal"
	"strings"
)

func RGB2Hex(data string) string {
	hits := internal.Regex("rgb", data)
	cleanVals := internal.GetCleanValues("rgb", hits)

	for i := 0; i < len(hits); i++ {

		numbers := strings.Split(cleanVals[i], ",")

		data = strings.Replace(data, hits[i], internal.CreateHex(numbers), -1)
	}

	return data
}
