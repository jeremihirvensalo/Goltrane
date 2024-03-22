package convert

var convFuncs = map[string]func(string) string{
	"hex2rgb":  Hex2RGB,
	"hex2rgba": Hex2RGBA,
	"rgb2hex":  RGB2Hex,
	"rgb2rgba": RGB2RGBA,
	"rgba2hex": RGBA2Hex,
}

func Call(method string, data string) string {
	return convFuncs[method](data)
}
