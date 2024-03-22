package internal

import "strconv"

func ToInt(str string) int {
	num, err := strconv.Atoi(str)
	CheckErr(err)

	return num
}
