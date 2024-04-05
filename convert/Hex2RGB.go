package convert

import (
	"fmt"
	"karhuhelsinki/goltrane/internal"
	"regexp"
	"strconv"
)

func Hex2RGB(data string) string {
	matches := internal.Regex("hex", data)

	for i := 0; i < len(matches); i++ {
		hex := matches[i][1:]

		hex = hex[:len(hex)-1]

		values, err := strconv.ParseUint(string(hex), 16, 32)
		internal.CheckErr(err)

		newVal := fmt.Sprintf("rgb(%d, %d, %d);", uint8(values>>16), uint8((values>>8)&0xFF), uint8(values&0xFF))

		reg := regexp.MustCompile(matches[i] + `{1}`)

		data = reg.ReplaceAllString(data, newVal)
	}

	return data
}
