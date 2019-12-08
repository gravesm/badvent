package exercises

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func Day2a(r io.Reader) int {
	c := NewComputer(r)
	return c.Compute(12, 2)
}

func Day2b(r io.Reader) int {
	c := NewComputer(r)
	var output int
	for a := 0; a < 100; a++ {
		for b := 0; b < 100; b++ {
			if c.Compute(a, b) == 19690720 {
				return 100*a + b
			}
		}
	}
	return output
}

type Computer struct {
	Program []int
	Tape    []int
}

func (c Computer) Compute(a, b int) int {
	var p int
	c.Reset()
	c.Tape[1] = a
	c.Tape[2] = b
Loop:
	for {
		out := c.Tape[p+3]
		x := c.Dereference(p + 1)
		y := c.Dereference(p + 2)
		switch c.Tape[p] {
		case 1:
			c.Tape[out] = x + y
		case 2:
			c.Tape[out] = x * y
		case 99:
			break Loop
		}
		p += 4
	}
	return c.Tape[0]
}

func (c *Computer) Reset() {
	tape := make([]int, len(c.Program))
	copy(tape, c.Program)
	c.Tape = tape
}

func (c Computer) Dereference(p int) int {
	return c.Tape[c.Tape[p]]
}

func NewComputer(r io.Reader) Computer {
	var ops []int
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		d := strings.Split(scanner.Text(), ",")
		for _, i := range d {
			op, _ := strconv.Atoi(i)
			ops = append(ops, op)
		}
	}
	return Computer{ops, []int{}}
}
