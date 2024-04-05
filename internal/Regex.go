package internal

import "regexp"

var RegexStrigns = map[string]string{
	"hex":         `(#[0-9a-f|A-F]{6}|#[0-9a-f|A-F]{3});`,
	"hexa":        `(#[0-9a-f|A-F]{8})`,
	"rgb":         `rgb\( *(\d{1,3} *, *\d{1,3} *, *\d{1,3}) *\)`,
	"rgba":        `rgba\( *(\d{1,3} *, *\d{1,3} *, *\d{1,3} *, *[0-1]?\.\d{1,2}?) *\)`,
	"hsl":         `hsl\( *(-?\d{1,3} *, *\d{1,3}%? *, *\d{1,3}%?) *\)`,
	"hsla":        `hsla\( *(\d{1,3} *, *\d{1,3}%? *, *\d{1,3}%? *, *[0-1]?\.\d{1,2}?) *\)`,
	"display-p3":  `color\( *display-p3 *([\.\d]* *[\.\d]* *[\.\d]* *) *\)`,
	"display-p3a": `color\( *display-p3 *([\.\d]* *[\.\d]* *[\.\d]* *\/ *[\.\d]* *) *\)`,
}

func Regex(colorType string, data string) []string {
	reg := regexp.MustCompile(RegexStrigns[colorType])

	return reg.FindAllString(string(data), -1)
}
