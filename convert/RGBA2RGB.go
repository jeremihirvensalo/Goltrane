package convert

import (
	"fmt"
	"karhuhelsinki/goltrane/internal"
	"math"
	"strings"
)

func RGBA2RGB(data string) string {
	hits := internal.Regex("rgba", data)
	cleanVals := internal.GetCleanValues("rgba", hits)

	for i := 0; i < len(cleanVals); i++ {
		numbers := strings.Split(cleanVals[i], ",")

		numbers[3] = internal.TransformAlpha(numbers[3])

		normRed := float64(internal.ToInt(numbers[0])) / 255.0
		normGreen := float64(internal.ToInt(numbers[1])) / 255.0
		normBlue := float64(internal.ToInt(numbers[2])) / 255.0

		alph := float64(internal.ToInt(numbers[3])) / 255.0

		red := int(math.Round(alph * normRed * 255.0))
		green := int(math.Round(alph * normGreen * 255.0))
		blue := int(math.Round(alph * normBlue * 255.0))

		rgb := fmt.Sprintf("rgb(%d, %d, %d)", red, green, blue)

		data = strings.Replace(data, hits[i], rgb, -1)
	}

	return data
}
