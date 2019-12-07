package main

import (
	"bufio"
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"strconv"
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
