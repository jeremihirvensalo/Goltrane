package convert

import (
	"fmt"
	"karhuhelsinki/goltrane/internal"
	"math"
	"strings"
)

func HSL2RGB(data string) string {
	hits := internal.Regex("hsl", data)
	cleanVals := internal.GetCleanValues("hsl", hits)

	for i := 0; i < len(hits); i++ {
		numbers := strings.Split(strings.Replace(cleanVals[i], "%", "", -1), ",")

		// The following mathematical equations are taken from https://www.baeldung.com/cs/convert-color-hsl-rgb

		hue := float64(internal.ToInt(numbers[0])) / 60.0  // Hue
		sat := float64(internal.ToInt(numbers[1])) / 100.0 // Saturation
		lig := float64(internal.ToInt(numbers[2])) / 100.0 // Lightness

		c := (1 - math.Abs(2*lig-1)) * sat
		x := c * (1 - math.Abs(math.Mod(hue, 2.0)-1))
		m := lig - (c / 2.0)

		r := int(math.Round(m * 255))       // Red
		g := int(math.Round((x + m) * 255)) // Green
		b := int(math.Round((c + m) * 255)) // Blue

		rgbColor := fmt.Sprintf("rgb(%d, %d, %d)", r, g, b)

		data = strings.Replace(data, hits[i], rgbColor, -1)

	}

	return data
}
