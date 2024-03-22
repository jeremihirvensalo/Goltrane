package convert

import (
	"fmt"
	"karhuhelsinki/goltrane/internal"
	"strconv"
	"strings"
)

func transformAlpha(alpha string) string {
	alphaVal, err := strconv.ParseFloat(alpha, 64)
	internal.CheckErr(err)

	// val := math.RoundToEven(alphaVal * 255)
	// val := alphaVal * 255
	return fmt.Sprintf("%d", int(alphaVal*255))

}

func RGBA2Hex(data string) string {

	hits := internal.Regex("rgba", data)
	cleanVals := internal.GetCleanValues("rgba", hits)

	for i := 0; i < len(cleanVals); i++ {
		numbers := strings.Split(cleanVals[i], ",")

		// alpha := numbers[3]

		// numbers = numbers[:3]

		// alpha, err := strconv.ParseFloat(numbers[3], 64)
		// internal.CheckErr(err)

		// alphaVal := int64(alpha * 255)

		numbers[3] = transformAlpha(numbers[3])

		hex := internal.CreateHex(numbers)

		data = strings.Replace(data, hits[i], hex, -1)
	}

	return data
}
