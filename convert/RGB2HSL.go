package convert

import (
	"karhuhelsinki/goltrane/internal"
	"strings"
)

func RGB2HSL(data string) string {

	hits := internal.Regex("rgb", data)
	cleanVals := internal.GetCleanValues("rgb", hits)

	for i := 0; i < len(hits); i++ {
		numbers := rgbValues(cleanVals[i])

		// Promising equation here https://www.niwa.nu/2013/05/math-behind-colorspace-conversions-rgb-hsl/

		lowest := findLowest(numbers)
		highest := findHighest(numbers)

		lum := (lowest + highest) / 2 // Lumiance

		// red := rgbVal(numbers[0])
		// green := rgbVal(numbers[1])
		// blue := rgbVal(numbers[2])

		// fmt.Println(red)
		// fmt.Println(green)
		// fmt.Println(blue)
	}

	return data
}

func rgbVal(str string) float64 {
	val := internal.ToInt(str)

	return float64(val) / 255
}

func rgbValues(str string) []float64 {
	values := strings.Split(str, ",")

	numbers := make([]float64, len(values))
	for i := 0; i < len(values); i++ {
		numbers[i] = rgbVal(values[i])
	}

	return numbers
}

func findLowest(numbers []float64) float64 {

	lowest := numbers[0]

	for _, val := range numbers {
		if val < lowest {
			lowest = val
		}
	}

	return lowest
}

func findHighest(numbers []float64) float64 {

	highest := numbers[0]
	for _, val := range numbers {
		if val > highest {
			highest = val
		}
	}

	return highest
}
