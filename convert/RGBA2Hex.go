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

	return fmt.Sprintf("%d", int(alphaVal*255))
}

func RGBA2Hex(data string) string {

	hits := internal.Regex("rgba", data)
	cleanVals := internal.GetCleanValues("rgba", hits)

	for i := 0; i < len(cleanVals); i++ {
		numbers := strings.Split(cleanVals[i], ",")

		numbers[3] = transformAlpha(numbers[3])

		hex := internal.CreateHex(numbers)

		data = strings.Replace(data, hits[i], hex, -1)
	}

	return data
}
