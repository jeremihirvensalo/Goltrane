package main

import (
	"karhuhelsinki/goltrane/command"

	"github.com/alecthomas/kong"
)

type CLI struct {
	Hex2RGB command.Hex2RGBCmd `cmd aliases:"hex2rgb" help:"Convert hex to rgb"`
}

func main() {
	cli := CLI{}
	ctx := kong.Parse(&cli)
	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}
