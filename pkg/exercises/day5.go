package exercises

import (
	"fmt"
	"github.com/gravesm/badvent/pkg/computer"
	"io"
)

func Day5(reader io.Reader) (int, int) {
	var one, two int
	var out1 []int
	p := computer.LoadProgram(reader)
	t := computer.NewTask()
	t.Load(p)
	go t.Run()
	t.In <- 1
	for output := range t.Out {
		out1 = append(out1, output)
	}
	for i := 0; i < len(out1)-1; i++ {
		fmt.Println(out1[i])
	}
	one = out1[len(out1)-1]
	t = computer.NewTask()
	t.Load(p)
	go t.Run()
	t.In <- 5
	two = <-t.Out
	return one, two
}
