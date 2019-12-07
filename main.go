package main

import (
	"bufio"
	"fmt"
	"github.com/urfave/cli/v2"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	app := &cli.App{
		Name:  "badvent",
		Usage: "run some bad code",
		Action: func(c *cli.Context) error {
			switch day := c.Args().Get(0); day {
			case "1.1":
				scanner := bufio.NewScanner(os.Stdin)
				var total int
				for scanner.Scan() {
					mass, _ := strconv.Atoi(scanner.Text())
					total += mass/3 - 2
				}
				fmt.Println(total)
			case "1.2":
				scanner := bufio.NewScanner(os.Stdin)
				var total int
				for scanner.Scan() {
					mass, _ := strconv.Atoi(scanner.Text())
					for fuel := mass/3 - 2; fuel > 0; fuel = fuel/3 - 2 {
						total += fuel
					}
				}
				fmt.Println(total)
			case "2.1":
				tape := make_tape(os.Stdin)
				output := compute(12, 2, tape)
				fmt.Println(output)
			case "2.2":
				tape := make_tape(os.Stdin)
				t := make([]int, len(tape))
			Loop:
				for n := 0; n < 100; n++ {
					for v := 0; v < 100; v++ {
						copy(t, tape)
						output := compute(n, v, t)
						if output == 19690720 {
							fmt.Println(100*n + v)
							break Loop
						}
					}
				}
			default:
				cli.ShowAppHelp(c)
			}
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func make_tape(r io.Reader) []int {
	reader := bufio.NewReader(r)
	input, _ := reader.ReadString('\n')
	codes := strings.Split(input[:len(input)-1], ",")
	tape := make([]int, len(codes))
	for idx, i := range codes {
		tape[idx], _ = strconv.Atoi(i)
	}
	return tape
}

func compute(n int, v int, tape []int) int {
	var pos int
	tape[1] = n
	tape[2] = v
Loop:
	for {
		code := tape[pos]
		out := tape[pos+3]
		x := tape[tape[pos+1]]
		y := tape[tape[pos+2]]
		switch code {
		case 1:
			tape[out] = x + y
		case 2:
			tape[out] = x * y
		case 99:
			break Loop
		}
		pos += 4
	}
	return tape[0]
}
