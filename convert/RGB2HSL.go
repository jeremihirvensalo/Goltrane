package convert

import (
	"karhuhelsinki/goltrane/internal"
)

func RGB2HSL(data string) string {

	hits := internal.Regex("rgb", data)
	cleanVals := internal.GetCleanValues("rba", hits)

	for i := 0; i < len(hits); i++ {

	}

	return data
}
