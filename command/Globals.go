package command

import (
	"bufio"
	"os"
)

type Globals struct {
	Infile  *os.File `name:"infile" short:"i" help:"Input file" default:"-"`
	Outfile string   `name:"outfile" short:"o" help:"Output file"`
}

/**
 * Reads input from file or stdin and returns the data as a string
 */
func (g *Globals) Read() (string, error) {
	var data []byte
	var err error

	data, err = bufio.NewReader(g.Infile).ReadBytes('\n')

	if err != nil {
		return "", err
	}
	return string(data), nil
}

/**
 * Writes data to file or stdout
 */
func (g *Globals) Write(data string) error {
	var err error

	if g.Outfile == "" {
		_, err = os.Stdout.Write([]byte(data))
	} else {
		err = os.WriteFile(g.Outfile, []byte(data), 0644)
	}

	return err
}
