package main

import (
	"fmt"
	"github.com/gravesm/badvent/pkg/exercises"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "badvent",
		Usage: "run some bad code",
		Action: func(c *cli.Context) error {
			switch day := c.Args().Get(0); day {
			case "1a":
				fmt.Println(exercises.Day1a(os.Stdin))
			case "1b":
				fmt.Println(exercises.Day1b(os.Stdin))
			case "2":
				a, b := exercises.Day2(os.Stdin)
				fmt.Printf("%d\n%d\n", a, b)
			case "3a":
				fmt.Println(exercises.Day3a(os.Stdin))
			case "3b":
				fmt.Println(exercises.Day3b(os.Stdin))
			case "4a":
				fmt.Println(exercises.Day4a(os.Stdin))
			case "4b":
				fmt.Println(exercises.Day4b(os.Stdin))
			case "5":
				a, b := exercises.Day5(os.Stdin)
				fmt.Printf("%d\n%d\n", a, b)
			case "6":
				a, b := exercises.Day6(os.Stdin)
				fmt.Printf("%d\n%d\n", a, b)
			case "7":
				a, b := exercises.Day7(os.Stdin)
				fmt.Printf("%d\n%d\n", a, b)
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
