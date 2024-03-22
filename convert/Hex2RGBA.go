package convert

import (
	"fmt"
	"karhuhelsinki/goltrane/internal"
	"regexp"
	"strconv"
)

func Hex2RGBA(data string) string {

	hits := internal.Regex("hex", data)

	for i := 0; i < len(hits); i++ {
		hex := hits[i][1:]

		values, err := strconv.ParseUint(string(hex), 16, 32)
		internal.CheckErr(err)

		newVal := fmt.Sprintf("rgba(%d, %d, %d, 1)", uint8(values>>16), uint8((values>>8)&0xFF), uint8(values&0xFF))

		reg := regexp.MustCompile(hits[i] + `{1}`)

		data = reg.ReplaceAllString(data, newVal)
	}

	return data
}
