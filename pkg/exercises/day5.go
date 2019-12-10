package exercises

import (
	"io"
)

func Day5a(r io.Reader) int {
	c := Computer{}
	c.Load(LoadProgram(r))
	return c.Run()
}
