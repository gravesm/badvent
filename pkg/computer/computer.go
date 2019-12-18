package computer

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Task struct {
	tape    []int
	counter int
	//instruction register
	reg_op int
	//argument registers
	reg_a1 int
	reg_a2 int
	reg_a3 int
	//return value register
	reg_rv *int
	//io buses
	Out chan int
	In  chan int
}

func NewTask() Task {
	return Task{Out: make(chan int, 2), In: make(chan int, 2)}
}

func (c *Task) Load(program []int) {
	tape := make([]int, len(program))
	copy(tape, program)
	c.tape = tape
}

func (c *Task) Run() {
	c.counter = 0
	for {
		c.decode()
		if !c.execute() {
			close(c.Out)
			close(c.In)
			break
		}
	}
}

func (c *Task) execute() bool {
	switch c.reg_op {
	case 1:
		*c.reg_rv = c.reg_a1 + c.reg_a2
	case 2:
		*c.reg_rv = c.reg_a1 * c.reg_a2
	case 3:
		*c.reg_rv = <-c.In
	case 4:
		c.Out <- c.reg_a1
	case 5:
		if c.reg_a1 != 0 {
			c.counter = c.reg_a2
		}
	case 6:
		if c.reg_a1 == 0 {
			c.counter = c.reg_a2
		}
	case 7:
		if c.reg_a1 < c.reg_a2 {
			*c.reg_rv = 1
		} else {
			*c.reg_rv = 0
		}
	case 8:
		if c.reg_a1 == c.reg_a2 {
			*c.reg_rv = 1
		} else {
			*c.reg_rv = 0
		}
	default:
		return false
	}
	return true
}

func (c *Task) decode() {
	inst := fmt.Sprintf("%05d", c.tape[c.counter])
	c.reg_op, _ = strconv.Atoi(inst[3:5])
	switch c.reg_op {
	case 1, 2, 7, 8:
		c.reg_rv = &c.tape[c.tape[c.counter+3]]
		c.reg_a1 = c.access(c.counter+1, int(inst[2]-48))
		c.reg_a2 = c.access(c.counter+2, int(inst[1]-48))
		c.counter += 4
	case 3:
		c.reg_rv = &c.tape[c.tape[c.counter+1]]
		c.counter += 2
	case 4:
		c.reg_a1 = c.access(c.counter+1, int(inst[2]-48))
		c.counter += 2
	case 5, 6:
		c.reg_a1 = c.access(c.counter+1, int(inst[2]-48))
		c.reg_a2 = c.access(c.counter+2, int(inst[1]-48))
		c.counter += 3
	}
}

func (c Task) access(p, m int) int {
	if m == 0 {
		return c.tape[c.tape[p]]
	}
	return c.tape[p]
}

func (c Task) Read(pos int) int {
	return c.tape[pos]
}

func LoadProgram(r io.Reader) []int {
	var ops []int
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		d := strings.Split(scanner.Text(), ",")
		for _, i := range d {
			op, _ := strconv.Atoi(i)
			ops = append(ops, op)
		}
	}
	return ops
}
