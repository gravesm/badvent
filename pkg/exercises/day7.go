package exercises

import (
	"github.com/gravesm/badvent/pkg/computer"
	"io"
)

func Day7(reader io.Reader) (int, int) {
	var max1, max2 int
	p := computer.LoadProgram(reader)
	for _, phases := range perm([]int{0, 1, 2, 3, 4}) {
		output := amp(phases, p)
		if output > max1 {
			max1 = output
		}
	}
	for _, phases := range perm([]int{5, 6, 7, 8, 9}) {
		output := amp(phases, p)
		if output > max2 {
			max2 = output
		}
	}
	return max1, max2
}

func amp(phases []int, program []int) int {
	var tasks []computer.Task
	for _, phase := range phases {
		t := computer.NewTask()
		t.Load(program)
		t.In <- phase
		tasks = append(tasks, t)
		go t.Run()
	}
	tasks[0].In <- 0
	for {
		select {
		case output, ok := <-tasks[0].Out:
			if !ok {
				tasks[0].Out = nil
			} else {
				tasks[1].In <- output
			}
		case output, ok := <-tasks[1].Out:
			if !ok {
				tasks[1].Out = nil
			} else {
				tasks[2].In <- output
			}
		case output, ok := <-tasks[2].Out:
			if !ok {
				tasks[2].Out = nil
			} else {
				tasks[3].In <- output
			}
		case output, ok := <-tasks[3].Out:
			if !ok {
				tasks[3].Out = nil
			} else {
				tasks[4].In <- output
			}
		case output := <-tasks[4].Out:
			tasks[0].In <- output
		}
		if tasks[0].Out == nil && tasks[1].Out == nil && tasks[2].Out == nil && tasks[3].Out == nil {
			output := <-tasks[4].Out
			return output
		}
	}
}

func perm(nums []int) [][]int {
	perms := [][]int{}
	var heaps func(int, []int)

	heaps = func(size int, nums []int) {
		if size == 1 {
			p := make([]int, len(nums))
			copy(p, nums)
			perms = append(perms, p)
		}
		for i := 0; i < size; i++ {
			heaps(size-1, nums)
			if size%2 == 0 {
				nums[i], nums[size-1] = nums[size-1], nums[i]
			} else {
				nums[0], nums[size-1] = nums[size-1], nums[0]
			}
		}
	}
	heaps(len(nums), nums)
	return perms
}
