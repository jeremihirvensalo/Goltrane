package convert

import (
	"encoding/hex"
	"fmt"
	"karhuhelsinki/goltrane/internal"
	"strconv"
	"strings"
)

func Hex2RGBA(data string) string {

	hits := internal.Regex("hexa", data)

	for i := 0; i < len(hits); i++ {
		hexVal := hits[i][1:]

		decodedHex, err := hex.DecodeString(hexVal)
		internal.CheckErr(err)

		alpha := float64(decodedHex[3]) / 255.0

		newColor := fmt.Sprintf("rgba(%d, %d, %d, %.2f)", decodedHex[0], decodedHex[1], decodedHex[2], alpha)

		data = strings.Replace(data, hits[i], newColor, -1)
	}

	// Translate hex values without alpha channel
	hexHits := internal.Regex("hex", data)
	for i := 0; i < len(hexHits); i++ {
		hexVal := hexHits[i][1:]

		values, err := strconv.ParseUint(string(hexVal), 16, 32)
		internal.CheckErr(err)

		newVal := fmt.Sprintf("rgba(%d, %d, %d, 1)", uint8(values>>16), uint8((values>>8)&0xFF), uint8(values&0xFF))

		data = strings.Replace(data, hexHits[i], newVal, -1)
	}

	return data
}
