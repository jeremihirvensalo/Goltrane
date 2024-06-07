package internal

import "fmt"

func CreateHex(numbers []string) string {
	hex := "#"

	for i := 0; i < len(numbers); i++ {
		hex += fmt.Sprintf("%02X", ToInt(numbers[i]))
	}

	return hex
}
