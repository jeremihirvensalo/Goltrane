package convert

import (
	"fmt"
	"karhuhelsinki/goltrane/internal"
	"math"
	"strings"
)

func RGB2HSL(data string) string {
	hits := internal.Regex("rgb", data)
	cleanVals := internal.GetCleanValues("rgb", hits)

	for i := 0; i < len(hits); i++ {
		numbers := strings.Split(cleanVals[i], ",")

		// The following mathematical equation is taken from https://www.niwa.nu/2013/05/math-behind-colorspace-conversions-rgb-hsl/

		// Convert the RGB values to the range 0 - 1
		red := float64(internal.ToInt(numbers[0])) / 255
		green := float64(internal.ToInt(numbers[1])) / 255
		blue := float64(internal.ToInt(numbers[2])) / 255

		// Find the maximum and minimum values of the RGB values
		max := math.Max(red, math.Max(green, blue))
		min := math.Min(red, math.Min(green, blue))

		lum := (min + max) / 2.0 // Luminance

		sat := max - min // Saturation
		if sat > 0 && lum <= 0.5 {
			sat = sat / (max + min)
		} else if sat > 0 && lum > 0.5 {
			sat = sat / (2.0 - sat)
		}

		hue := (max - min) // Hue

		if hue > 0 {
			if red == max {
				hue = (green - blue) / hue
			} else if green == max {
				hue = 2.0 + (blue-red)/hue
			} else {
				hue = 4.0 + (red-green)/hue
			}
		}

		// Convert hue to degrees
		hue *= 60
		if hue < 0 {
			hue += 360
		}

		// Round the values to the nearest integer
		h := int(math.Round(hue))
		s := int(math.Round(sat * 100.0))
		l := int(math.Round(lum * 100.0))

		hsl := fmt.Sprintf("hsl(%d, %d, %d)", h, s, l)

		data = strings.Replace(data, hits[i], hsl, -1)
	}

	return data
}
