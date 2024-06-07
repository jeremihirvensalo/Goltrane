package convert

import (
	"karhuhelsinki/goltrane/internal"
	"strings"
)

func RGBA2RGB(data string) string {
	hits := internal.Regex("rgba", data)
	cleanVals := internal.GetCleanValues("rgba", hits)

	for i := 0; i < len(cleanVals); i++ {
		numbers := strings.Split(cleanVals[i], ",")

		numbers[3] = internal.TransformAlpha(numbers[3])
	}

	return data
}
