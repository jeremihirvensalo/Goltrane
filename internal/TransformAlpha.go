package internal

import (
	"fmt"
	"strconv"
)

func TransformAlpha(alpha string) string {
	alphaVal, err := strconv.ParseFloat(alpha, 64)
	CheckErr(err)

	return fmt.Sprintf("%d", int(alphaVal*255))
}
