package exercises

import (
	"bufio"
	"io"
	"strconv"
)

func Day1a(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	var total int
	for scanner.Scan() {
		mass, _ := strconv.Atoi(scanner.Text())
		total += mass/3 - 2
	}
	return total
}

func Day1b(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	var total int
	for scanner.Scan() {
		mass, _ := strconv.Atoi(scanner.Text())
		for f := mass/3 - 2; f > 0; f = f/3 - 2 {
			total += f
		}
	}
	return total
}
