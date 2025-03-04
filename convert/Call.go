package convert

var convFuncs = map[string]func(string) string{
	"hex2rgb":  Hex2RGB,
	"hex2rgba": Hex2RGBA,
	"rgb2hex":  RGB2Hex,
	"rgb2rgba": RGB2RGBA,
	"rgb2hsl":  RGB2HSL,
	"rgba2hex": RGBA2Hex,
	"hsl2rgb":  HSL2RGB,
}

func Call(method string, data string) string {
	return convFuncs[method](data)
}
