package command

import (
	"karhuhelsinki/goltrane/convert"
	"os"
)

type Hex2RGBCmd struct {
	Input  string `name:"infile" help:"Input file"`
	Output string `name:"outfile" help:"Output file"`
}

func (cmd *Hex2RGBCmd) Run() error {
	data, err := os.ReadFile(cmd.Input)

	if err != nil {
		return err
	}

	translated := convert.Hex2RGB(string(data))
	err = os.WriteFile(cmd.Output, []byte(translated), 0644)

	return err
}
