package command

import (
	"karhuhelsinki/goltrane/convert"
)

type Hex2RGBCmd struct {
	Globals
}

func (cmd *Hex2RGBCmd) Run() error {
	data, err := cmd.Read()

	if err != nil {
		return err
	}

	translated := convert.Hex2RGB(data)

	return cmd.Write(translated)
}
