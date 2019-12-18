package exercises

import (
	"github.com/gravesm/badvent/pkg/computer"
	"io"
)

func Day2(reader io.Reader) (int, int) {
	var one, two int
	p := computer.LoadProgram(reader)
	t := computer.NewTask()
	p[1], p[2] = 12, 2
	t.Load(p)
	t.Run()
	one = t.Read(0)
Loop:
	for a := 0; a < 100; a++ {
		for b := 0; b < 100; b++ {
			t := computer.NewTask()
			p[1], p[2] = a, b
			t.Load(p)
			t.Run()
			if t.Read(0) == 19690720 {
				two = 100*a + b
				break Loop
			}
		}
	}
	return one, two
}
